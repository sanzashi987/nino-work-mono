import React, { useMemo } from 'react';
import { useSpringBasic } from '@canvas/runtime-components';
import { animated } from '@react-spring/web';
import styles from './container.module.scss';
import type { RuntimeInterface } from '../types';
import type { ResponsiveController } from '@/component/Controller';
import { HiddenMode } from '@/types';

const previewClass = `${styles['preview-box']} canvix-container__reserved`;

export type PreviewStaticMinimalProps = {
  containerRef: React.RefObject<HTMLDivElement | null>;
  transitionRef: React.RefObject<any>;
  children: React.ReactNode;
  basic: Record<string, any>;
  projectId: string;
  className?: string;
  id?: string;
} & Pick<ResponsiveController.Props['config'], 'type' | 'hide'>;

function createContainer({ getRealBasic }: RuntimeInterface) {
  const PreviewContainerMinimal: React.FC<PreviewStaticMinimalProps> = ({
    containerRef,
    transitionRef,
    children,
    type,
    hide,
    basic,
    projectId,
    id,
    className
  }) => {
    const { springBasic, normalBasic } = useMemo(
      () => getRealBasic(basic, projectId),
      [basic, projectId]
    );

    const style = { ...useSpringBasic(springBasic, !!hide, transitionRef), ...normalBasic };
    if (hide === HiddenMode.unmount) {
      containerRef.current = null;
      return null;
    }
    if (type === 'subcom') {
      return children;
    }

    return (
      <animated.div
        id={id}
        ref={containerRef}
        className={className ?? previewClass}
        style={style}
      >
        {children}
      </animated.div>
    );
  };

  const PreviewContainer: React.FC<ResponsiveController.ContainerProps> = ({
    runtime: { config, utils },
    transitionRef,
    children,
    outer
  }) => {
    const { type, hide } = config;
    const basicEnsured = useMemo(() => (config.type === 'subcom' ? {} : config.basic), [config]);

    return (
      <PreviewContainerMinimal
        basic={basicEnsured}
        type={type}
        hide={hide}
        containerRef={utils.containerRef}
        transitionRef={transitionRef}
        projectId={outer.projectId}
      >
        {children}
      </PreviewContainerMinimal>
    );
  };
  return {
    PreviewContainer,
    PreviewContainerMinimal
  };
}

export { createContainer };
