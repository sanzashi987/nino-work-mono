import React, { ComponentType } from 'react';
import sandbox from '@/component/sandbox';
import { connect } from './connector';
import { createComponentLoader } from '../ComponentLoader';
import { createContainer } from './Static';
import Controller, {
  BasicStates, FullUtils, ResponsiveController, ResponsivePanelUtils, ResponsivePanelUtilsInsideWrapper, ServiceConnector
} from '@/component/Controller';

import { typeToService, ServiceComponent } from '@/component/services';
import { ConfigTypeSupportedInControllerRuntime, LayerList, ConnectorProps, isUnmountMode } from '@/types';
import { createUtils } from './utils';
import { ComWrapperProps, RuntimeInterface } from '../../types/com-config/connector';
import RuntimeError from '@/component/RuntimeError';

abstract class ComWrapperType extends Controller<
ConfigTypeSupportedInControllerRuntime,
ResponsivePanelUtils,
LayerList,
ResponsiveController.OptionProps,
ResponsivePanelUtilsInsideWrapper
> {
  declare Container: ComponentType<ResponsiveController.ContainerProps>;
  // declare switchRenderer: (props: ComWrapperProps) => void;

  declare createUtils: () => FullUtils<ResponsivePanelUtilsInsideWrapper>;

  declare mounted: () => void;
}

interface ComWrapperClass {
  new (props: ComWrapperProps): ComWrapperType;
}

const servicesKeptInUnmount = Object.entries(typeToService)
  .filter((val) => val[0] === 'com')
  .map((val) => val[1]);

export function createComponentWrapper(runtimeImpl: RuntimeInterface) {
  const ComponentLoader = createComponentLoader(runtimeImpl);
  const { PreviewContainer, PreviewContainerMinimal } = createContainer(runtimeImpl);
  class ComWrapper extends Controller<ConfigTypeSupportedInControllerRuntime,
  ResponsivePanelUtils,
  LayerList,
  ResponsiveController.OptionProps,
  ResponsivePanelUtilsInsideWrapper
  > {
    isError = false;

    Container: ComponentType<ResponsiveController.ContainerProps> = PreviewContainer;

    constructor(props: ComWrapper['props']) {
      super(props);
      if (isUnmountMode(props.config.hide)) {
        this.runtimeServices = this.getServicesKeptInUnmount();
      }
    }

    getServicesKeptInUnmount() {
      return servicesKeptInUnmount;
    }

    shouldComponentUpdate(nextProps: ComWrapper['props'], nextState: any) {
      const shouldRender = super.shouldComponentUpdate(nextProps, nextState);
      return shouldRender || nextProps.chain !== this.props.chain;
    }

    componentDidCatch(e: any, info: any) {
      console.error(e, info);
      // this.initState = {};
      this.errorForce(true);
    }

    componentDidMount(): void {
      if (isUnmountMode(this.props.config.hide) && !this._ready) {
        this._ready = true;
        this.core.start();
      }
    }

    componentDidUpdate(prevProps: ComWrapper['props'], prevState: ComWrapper['state']): void {
      const { hide } = this.state.config;
      if (hide !== prevState.config.hide) {
        // switch to unmount mode
        if (isUnmountMode(hide)) {
          this.runtimeServices = this.getServicesKeptInUnmount();
          this.cleanup();
          this.forceUpdate();
          // switch from unmount mode
        } else if (isUnmountMode(prevState.config.hide)) {
          this.runtimeServices = this.usedServices;
          this.emitInstance();
          this.forceUpdate();
        }
      }
    }

    mounted(): void {
      if (!this._ready) {
        this._ready = true;
        this.core.start();
        this.emitInstance();
      }
    }

    errorForce = (error: boolean) => {
      this.isError = error;
      this.forceUpdate();
      this.props.forceUpdate();
    };

    retry = () => {
      this.errorForce(false);
    };

    createUtils(): FullUtils<ResponsivePanelUtilsInsideWrapper> {
      return createUtils(this, runtimeImpl.getAssetsUrl);
    }

    render(): React.ReactNode {
      const runtimeProps = this.getRuntimeProps();
      const { Container } = this;
      return (
        <>
          {this.renderServices()}
          <Container
            ready={this._ready}
            outer={this.props}
            runtime={runtimeProps}
            transitionRef={this.transitionRef}
          >
            {this.isError ? (
              <RuntimeError name={this.props.config?.name} retry={this.retry} />
            ) : (
              <ComponentLoader {...runtimeProps} chain={this.props.chain} />
            )}
          </Container>
        </>
      );
    }
  }

  const PreviewWrapper = connect(ComWrapper);

  let nodeServices: Record<string, ServiceComponent> | null = null;

  class PreviewNode extends ServiceConnector<
  LogicalNodeProps,
  BasicStates<LogicalNodeProps['config']>
  > {
    constantProps;

    constructor(props: ConnectorProps<LogicalNodeProps>) {
      super(props);
      this.constantProps = {
        ref: this.ref,
        mounted: this.mounted.bind(this),
        utils: {
          ...props.logicalUtils,
          $emit: this.emit
        }
      };
    }

    getEnabledServices(): Record<string, ServiceComponent> {
      if (nodeServices) return nodeServices;
      nodeServices = { ...typeToService };
      delete nodeServices.basic;
      delete nodeServices.attr;
      return nodeServices;
    }

    getIdentifier() {
      return {
        ...this.props.config.com,
        panelId: this.props.panelId,
        dashboardId: this.props.projectId,
        comId: this.props.config.id
      };
    }

    mounted() {
      this.core.start();
    }

    render() {
      const runtimeConfig = getRuntimeConfig({
        input: this.props.config,
        runtimeKeys: ['attr'],
        config: {}
      });

      return (
        <>
          {this.renderServices()}
          <ComponentLoader
            config={runtimeConfig}
            data={this.state.data}
            {...this.constantProps}
            sandboxRunner={sandbox.runInSandbox}
          />
        </>
      );
    }
  }

  return {
    ComWrapper: ComWrapper as unknown as ComWrapperClass,
    PreviewWrapper: PreviewWrapper as ConnectOuptut,
    PreviewNode: PreviewNode as NodeComType,
    PreviewContainer,
    PreviewContainerMinimal
  };
}
