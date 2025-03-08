import React, { useMemo } from 'react';
import { Responsive, type ResponsiveController } from '@canvas/component-factory';
import { useSpringBasic } from '@canvas/runtime-components';
import { animated } from 'react-spring';

import styles from './container.module.scss';
import type { RuntimeInterface } from '../context';

const CONTAINER_FINDER = 'canvas-container__reserved';
const previewClass = styles['preview-box'] + ' ' + CONTAINER_FINDER;

export type PreviewStaticMinimalProps = {
  containerRef: React.MutableRefObject<HTMLDivElement | null>;
  transitionRef: React.MutableRefObject<any>;
  children: React.ReactNode;
  basic: Record<string, any>;
  dashboardId: string;
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
    dashboardId,
    id,
    className,
  }) => {
    // const _basic = useMemo(() => {
    //   return type === 'subcom' ? {} : basic;
    // }, [type, basic]);

    const { springBasic, normalBasic } = useMemo(
      () => getRealBasic(basic, dashboardId),
      [basic, dashboardId],
    );

    const style = { ...useSpringBasic(springBasic, !!hide, transitionRef), ...normalBasic };
    if (hide === Responsive.HiddenMode.nonexistent) {
      containerRef.current = null;
      return null;
    }
    if (type === 'subcom') {
      return <>{children}</>;
    }

    return (
      <animated.div
        id={id ?? ''}
        ref={containerRef}
        className={className ?? previewClass}
        //@ts-ignore
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
    outer,
  }) => {
    const { type, hide } = config;
    const basicEnsured = useMemo(() => {
      return config.type === 'subcom' ? {} : config.basic;
    }, [config]);

    return (
      <PreviewContainerMinimal
        basic={basicEnsured}
        type={type}
        hide={hide}
        containerRef={utils.containerRef}
        transitionRef={transitionRef}
        dashboardId={outer.dashboardId}
      >
        {children}
      </PreviewContainerMinimal>
    );
  };
  return {
    PreviewContainer,
    PreviewContainerMinimal,
  };
}

export { createContainer };
