import React, {
  Component, lazy, Suspense, ComponentType, ReactNode
} from 'react';

type LoaderType = Parameters<typeof lazy>[0];

type LoadableProps = {
  loader: LoaderType;
  loading: () => NonNullable<ReactNode> | null;
};

export function Loadable({ loader, loading }: LoadableProps): ComponentType<any> {
  const Com = lazy(loader);
  // eslint-disable-next-line react/prefer-stateless-function
  return class LazyLoad extends Component {
    render() {
      return (
        <Suspense fallback={loading()}>
          <Com {...this.props} />
        </Suspense>
      );
    }
  };
}

const loading = () => null;

export const BlankLoadable = (loader: LoaderType) => Loadable({ loader, loading });
