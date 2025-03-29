/* eslint-disable @typescript-eslint/no-namespace */
declare global {
  namespace NodeJS {
    interface ProcessEnv {
      readonly NODE_ENV: 'development' | 'production' | 'test';
    }
  }
}

export const NINO_IS_PROD = process.env.NODE_ENV === 'production';
export const NINO_IS_DEV = process.env.NODE_ENV === 'development';
export const NINO_IS_TEST = process.env.NODE_ENV === 'test';
