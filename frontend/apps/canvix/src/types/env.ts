type SPAEnv = {
  platform: 'SPA';
};

type H5Env = {
  platform: 'H5';
};

type NativeEnv = {
  instanceId: string;
  platform: 'IOS' | 'Android';
  bridgeId: string;
  channel: {
    postMessage: (payload: string) => void;
  };
};

type GetPlatformEnv<T> = Extract<SPAEnv | H5Env | NativeEnv, { platform: T }>;

export type EnvVariables<T> = {
  projectId: string;
} & GetPlatformEnv<T>;
