import { queryStringify } from '@canvix/utils';
import { CacheComponentType, LoadStatus } from './type';

class Dub {
  url: string;

  qsParams: Record<string, any>;

  cacheLoad: Record<string, CacheComponentType>;

  componentDeps: Record<string, string[]>;

  dependeds: Record<string, string[]>;

  priorities: Record<string, string[]>;

  recordScopeModules: Record<string, string>;

  currentModule: string;

  scopeModule: string;

  constructor(url: string, qsParams?: Record<string, any>) {
    this.url = url;
    this.qsParams = qsParams || {};
    this.cacheLoad = {};
    this.componentDeps = {};
    this.dependeds = {};
    this.priorities = {};
    this.recordScopeModules = {};
    this.currentModule = '';
    this.scopeModule = '';
  }

  loadModule(moduleName: string, deps: string[], callback: (module: any) => void): void;
  loadModule(moduleName: string, deps: string[], order: string[], callback: (module: any) => void): void;
  loadModule(...args: any[]) {
    let name = args[0];
    const callback = args[args.length - 1];
    let deps: string[] = [];
    let orders: string[] = [];
    if (args.length > 2) {
      deps = args.at(1)!;
    }
    if (args.length > 3) {
      orders = args.at(2);
    }
    this.getNameSpaces(deps);
    name = this.modifyCustomName(name);
    deps = this.parseCustomComDeps(name, deps);
    if (!this.cacheLoad[name]) {
      this.initModuleInfo(name);
    }
    this.cacheLoad[name].factory = callback;
    this.cacheLoad[name].loaded = LoadStatus.RESOLVED;
    this.priorities[name] = orders;
    this.handleModule(name, deps);
  }

  handleModule = (name: string, deps: string[]) => {
    const { componentDeps, cacheLoad, dependeds } = this;
    const isAllDependenicesFire = this.checkDependenciesFired(name, deps);
    componentDeps[name] = deps;
    if (isAllDependenicesFire) {
      const loadedComponent = cacheLoad[name];
      const params = this.getDependenices(name);
      const result = loadedComponent.factory?.(...params);
      loadedComponent.fired = true;
      if (result) {
        loadedComponent.exports = result;
      }
      if (dependeds[name]) {
        dependeds[name].filter((m) => cacheLoad[m] && !cacheLoad[m].fired).forEach((m) => {
          this.handleModule(m, componentDeps[m]);
        });
      }
    }
  };

  getDependenices = (name: string) => {
    const { componentDeps, cacheLoad } = this;
    return componentDeps[name] ? componentDeps[name].map((m: string) => cacheLoad[m].exports) : [];
  };

  initModuleInfo = (name: string) => {
    this.cacheLoad[name] = { exports: {}, loaded: LoadStatus.NULL, fired: false };
  };

  getNameSpaces = (depends: string[]) => {
    const scopeTagReg = /\/mobile\//;
    if (depends.length === 1 && scopeTagReg.test(depends[0])) {
      const [user, module] = depends[0].split(scopeTagReg);
      this.currentModule = `mobile/${module}`;
      this.scopeModule = user;
      this.recordScopeModules[this.currentModule] = depends.at(0)!;
    }
  };

  parseCustomComDeps = (name: string, deps: string[]) => {
    const { scopeModule, recordScopeModules } = this;
    if (scopeModule === name) {
      deps = deps.map((item) => {
        if (/^npm\//.test(item)) {
          recordScopeModules[item] = `${scopeModule}/${item}`;
          return recordScopeModules[item];
        }
        return item;
      });
    }
    return deps;
  };

  initScopeVars = () => {
    this.recordScopeModules = {};
    this.scopeModule = '';
    this.currentModule = '';
  };

  modifyCustomName = (name: string) => {
    const { recordScopeModules } = this;
    if (recordScopeModules[name]) {
      name = recordScopeModules[name];
      delete recordScopeModules[name];
    }
    if (Object.keys(recordScopeModules).length === 0) {
      this.initScopeVars();
    }
    return name;
  };

  checkHighPriority = (name: string, module: string) => {
    const { priorities, cacheLoad } = this;
    const parallelDeps = priorities[name];
    const index = parallelDeps.findIndex((m: string) => m === module);
    if (index <= 0) return false;
    return parallelDeps
      .slice(0, index)
      .some((pkg: string) => !cacheLoad[pkg] || cacheLoad[pkg].loaded !== LoadStatus.RESOLVED);
  };

  checkDependenciesFired = (name: string, deps: string[] = []) => {
    const { cacheLoad, dependeds, fetchRemoteComponent } = this;
    let flag = true;
    for (let i = 0; i < deps.length; i++) {
      const m = deps[i];
      if (!cacheLoad[m]) {
        this.initModuleInfo(m);
      }
      if (m === 'exports') {
        cacheLoad[m] = {
          exports: cacheLoad[name].exports,
          loaded: LoadStatus.RESOLVED,
          fired: true
        };
        continue;
      }
      if (!(cacheLoad[m] && cacheLoad[m].fired) && !this.checkHighPriority(name, m)) {
        flag = false;
        if (cacheLoad[m]?.loaded !== LoadStatus.PENDING) {
          fetchRemoteComponent(m);
        }
      }
      if (!dependeds[m]) {
        dependeds[m] = [];
      }
      if (!dependeds[m].includes(name)) {
        dependeds[m].push(name);
      }
    }
    return flag;
  };

  fetchRemoteComponent = (m: string) => {
    const { url: staticBaseUrl, cacheLoad, qsParams } = this;
    const element: HTMLScriptElement = document.createElement('script');
    const url = `${staticBaseUrl}/${m}/index.js?${queryStringify(qsParams)}`;
    element.src = url;
    element.type = 'text/javascript';
    element.charset = 'utf-8';
    element.async = true;
    cacheLoad[m].loaded = LoadStatus.PENDING;
    const loadCallback = () => {
      element.removeEventListener('load', loadCallback);
    };
    const errorCallback = (e: ErrorEvent) => {
      const error = e.error || new Error(`Load Component failed. src=${url}`);
      element.removeEventListener('error', errorCallback);
      console.log(error);
    };
    element.addEventListener('load', loadCallback);
    element.addEventListener('error', errorCallback);
    document.body.appendChild(element);
  };
}
const dubCreator = (url: string, qsParams?: Record<string, any>) => {
  const duber = new Dub(url, qsParams);
  return duber.loadModule;
};
export default dubCreator;
export { composeCacheId } from './utils';
