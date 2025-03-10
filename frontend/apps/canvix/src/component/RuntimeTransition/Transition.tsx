/* eslint-disable no-nested-ternary */
import React, {
  useState,
  forwardRef,
  useImperativeHandle,
  ReactNode,
  useRef,
  CSSProperties
} from 'react';
import { CSSTransition } from 'react-transition-group';
import './animate.css';
import styles from './index.module.scss';

const { hide, show } = styles;
const animatedHide = `animate__animated ${hide}`;
const animatedShow = `animate__animated ${show}`;

type TransitionEnhancerProps = {
  display: boolean;
  opacity: number;
  children: (style: CSSProperties) => ReactNode;
};

// directy pass props like `display` and `opacity` to `div` element may leads to a native type error
type TransitionConfig = {
  display: boolean;
  duration: number;
  opacity: number;
};

type TransitionInputType = {
  type: string;
  duration: number;
};

const defaultEnter = { type: 'fadeIn', duration: 0 };
const defaultExit = { type: 'fadeOut', duration: 0 };

function provideStyle({ opacity, duration, display }: TransitionConfig, isAnimated = false) {
  const res = { opacity };
  return Object.assign(
    res,
    !isAnimated
      ? display
        ? {}
        : { display: 'none' }
      : { '--animate-duration': `${duration}ms`, '--animate-opacity': `${opacity}` }
  );
}

function ensurePrefix(className: any) {
  if (typeof className !== 'string') return '';
  return className.startsWith('animate__') ? className : `animate__${className}`;
}

function ensureEntry(input: any, defaultConfig: TransitionInputType): TransitionInputType {
  const isObject = !!input && typeof input === 'object' && !(input instanceof Array);
  if (isObject) {
    return {
      type: input.type ?? defaultConfig.type,
      duration: input.duration ?? defaultConfig.duration
    };
  }
  return defaultConfig;
}

const Transition = forwardRef(
  ({ display, opacity, children }: TransitionEnhancerProps, ref: any) => {
    const isAnimated = useRef<boolean>(false);
    const [enterTransition, setEnterTransition] = useState<TransitionInputType>(defaultEnter);
    const [exitTransition, setExitTransition] = useState<TransitionInputType>(defaultExit);
    useImperativeHandle(
      ref,
      () => ({
        display: (config?: TransitionInputType) => {
          isAnimated.current = true;
          const validated = ensureEntry(config, defaultEnter);
          setEnterTransition(validated);
        },
        hide: (config?: TransitionInputType) => {
          isAnimated.current = true;
          const validated = ensureEntry(config, defaultExit);
          setExitTransition(validated);
        }
      }),
      []
    );

    const { duration } = display ? enterTransition : exitTransition;
    const enterActive = ensurePrefix(enterTransition.type);
    const exitActive = ensurePrefix(exitTransition.type);
    const style = provideStyle({ opacity, duration, display }, isAnimated.current);
    // const classNames = provideClassNames();
    return (
      <CSSTransition
        in={display}
        classNames={{
          enter: animatedHide,
          enterActive: `${enterActive} ${show}`,
          enterDone: animatedShow,
          exit: animatedShow,
          exitActive: `${exitActive} ${show}`,
          exitDone: animatedHide
        }}
        timeout={{
          enter: enterTransition.duration,
          exit: exitTransition.duration
        }}
      >
        {children(style)}
      </CSSTransition>
    );

    // return renderProps(style)
  }
);

Transition.displayName = 'TransitionEnhancer';

export default Transition;
