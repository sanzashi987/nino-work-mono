import React from 'react';
import ReactDOM from 'react-dom/client';
import type { BaseConfig, BaseModalProps, BaseConfigRuntime, ModalComponent } from './types';

type CollectionMeta = {
  root: ReactDOM.Root;
  container: HTMLDivElement;
  visible: boolean;
  config: BaseConfig;
  onClose: BaseModalProps['onClose'];
};
const collections: Record<string, CollectionMeta> = {};

function destroy(id: string) {
  const modal = collections[id];
  if (!modal?.root || !modal.container) return;
  const { root, container } = modal;
  // unmount component in async way to avoid error
  setTimeout(() => {
    root.unmount();
  });
  if (container.parentNode) {
    delete collections[id];
    container.parentNode.removeChild(container);
  }
}

const createContainer = (id: string) => {
  const modal = collections[id];
  if (modal?.root && modal?.container) {
    return { root: modal.root, container: modal.container };
  }
  const container = document.createElement('div');
  document.body.appendChild(container);
  const root = ReactDOM.createRoot(container!);
  return { root, container };
};

function activate(Component: ModalComponent, config: BaseConfig) {
  let options: BaseConfigRuntime = { ...config, visible: true };

  function onClose(...args: any) {
    const modal = collections[options.id];
    if (!modal?.visible) return;
    /**
     * @params args
     * @see {@link BaseConfig onClose}
     */
    if (args[0] === 'backdropClick' && options.disableBackdropClick) return;
    options.visible = false;
    render(options);
    options.onClose?.(...args);
  }

  function afterClose() {
    if (!options.keepMounted) {
      destroy(options.id);
    }
    options.afterClose?.();
  }

  const render = (props: BaseConfigRuntime) => {
    const modal = collections[props.id];
    if (!modal?.root) return;
    const { visible, content, ...other } = props;
    modal.config = { ...other, content };
    modal.visible = visible;
    const params = { ...other, visible, onClose, afterClose };
    modal.root.render(<Component {...params}>{content}</Component>);
  };

  function update(conf: Omit<BaseConfig, 'id'>) {
    options = { ...options, ...conf };
    render(options);
  }

  const modal = collections[options.id];
  const modalContent = options.content || modal?.config?.content;
  options.content = modalContent;
  const { visible, ...otherConfig } = options;
  const wrapper = createContainer(options.id);
  collections[options.id] = { ...wrapper, config: otherConfig, visible, onClose };
  render(options);

  return { id: options.id, close: onClose, update };
}

export default activate;
export { collections };
