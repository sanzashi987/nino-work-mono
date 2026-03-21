import type { Configuration as RspackConfiguration } from '@rspack/core';
import type { RsbuildConfig } from '@rsbuild/core';
import type { Config as TailwindConfig } from 'tailwindcss';

/**
 * Infrastructure configuration for nino-work monorepo apps
 *
 * @example
 * ```javascript
 * // infra.config.js
 * const path = require('path');
 *
 * module.exports = {
 *   port: 3000,
 *   mode: 'micro-host',
 *   devServer(config) {
 *     config.proxy = {
 *       '/backend': {
 *         target: 'http://localhost:8081',
 *         changeOrigin: true,
 *       },
 *     };
 *   },
 * };
 * ```
 */
export interface InfraConfig {
  /**
   * Development server port
   * @default 3000
   */
  port?: number;

  /**
   * Application running mode
   * - `null` or `undefined`: Standalone mode - independent app with HTML entry
   * - `'micro-host'`: Micro-frontend host app - loads child apps via Module Federation
   * - `'micro-app'`: Micro-frontend child app - exposed via Module Federation, no HTML generated
   *
   * Can be overridden via CLI: `pnpm start-mf --mode=micro-host`
   *
   * @default null
   */
  mode?: null | 'micro-host' | 'micro-app';

  /**
   * Tailwind CSS configuration
   * Passed directly to @tailwindcss/postcss plugin
   *
   * Requires: `pnpm add -D tailwindcss @tailwindcss/postcss`
   *
   * @example
   * ```javascript
   * tailwind: {
   *   content: ['./src/**\/*.{html,js,ts,jsx,tsx}'],
   *   theme: {
   *     extend: {},
   *   },
   * }
   * ```
   */
  tailwind?: TailwindConfig;

  /**
   * Global constants defined at build time
   * These are replaced during bundling using Rsbuild's define feature
   *
   * @example
   * ```javascript
   * define: {
   *   __VERSION__: JSON.stringify('1.0.0'),
   *   __DEV__: JSON.stringify(true),
   * }
   * ```
   */
  define?: Record<string, string>;

  /**
   * Customize development server configuration
   * Called with the dev server config object for modification
   *
   * @param config - DevServer configuration object to mutate
   *
   * @example
   * ```javascript
   * devServer(config) {
   *   config.proxy = {
   *     '/backend': {
   *       target: 'http://localhost:8081',
   *       changeOrigin: true,
   *     },
   *   };
   *   config.headers = {
   *     'X-Custom-Header': 'value',
   *   };
   * }
   * ```
   */
  devServer?(config: DevServerConfig): void;

  /**
   * Customize Rspack configuration
   * Similar to webpack configuration, allows deep customization
   *
   * @param config - Rspack configuration object
   * @returns Modified configuration (must return the config)
   *
   * @example
   * ```javascript
   * webpack(config) {
   *   config.resolve.alias['@'] = path.resolve(__dirname, 'src');
   *   config.plugins.push(new MyPlugin());
   *   return config;
   * }
   * ```
   */
  webpack?(config: RspackConfiguration): RspackConfiguration | undefined;
}

/**
 * Development server configuration
 * Based on Rsbuild's DevServerConfig
 */
export interface DevServerConfig {
  /**
   * Dev server port
   */
  port?: number;

  /**
   * Host to listen on
   * @default '0.0.0.0'
   */
  host?: string;

  /**
   * Enable Hot Module Replacement
   * @default true
   */
  hot?: boolean;

  /**
   * Proxy configuration for API requests
   *
   * @example
   * ```javascript
   * proxy: {
   *   '/api': {
   *     target: 'http://localhost:8080',
   *     changeOrigin: true,
   *     pathRewrite: { '^/api': '' },
   *   },
   * }
   * ```
   */
  proxy?: Record<string, ProxyConfig>;

  /**
   * Custom headers
   */
  headers?: Record<string, string>;

  /**
   * Enable history API fallback for SPA routing
   * @default true
   */
  historyApiFallback?: boolean;

  /**
   * Enable HTTPS
   */
  https?: boolean;

  /**
   * Custom dev server options
   */
  [key: string]: any;
}

/**
 * Proxy configuration for dev server
 */
export interface ProxyConfig {
  /**
   * Target server URL
   */
  target: string;

  /**
   * Change origin header to target URL
   * @default false
   */
  changeOrigin?: boolean;

  /**
   * Rewrite URL paths
   *
   * @example
   * ```javascript
   * pathRewrite: { '^/api': '' }
   * ```
   */
  pathRewrite?: Record<string, string>;

  /**
   * Custom headers for proxied requests
   */
  headers?: Record<string, string>;

  /**
   * Enable WebSocket proxying
   */
  ws?: boolean;

  /**
   * Only proxy requests matching this path
   */
  context?: string | string[] | ((pathname: string) => boolean);

  /**
   * Additional proxy options
   */
  [key: string]: any;
}

/**
 * Module Federation configuration (auto-generated for micro-app mode)
 */
export interface ModuleFederationConfig {
  /**
   * Name of the remote module
   * Auto-generated from package.json name (non-alphanumeric chars replaced with _)
   */
  name: string;

  /**
   * Modules to expose
   * @default {}
   */
  exposes?: Record<string, string>;

  /**
   * Shared dependencies
   * @default { react, react-dom, single-spa }
   */
  shared?: Record<string, SharedDependency>;
}

/**
 * Shared dependency configuration for Module Federation
 */
export interface SharedDependency {
  /**
   * Use singleton instance
   */
  singleton?: boolean;

  /**
   * Load dependency eagerly
   */
  eager?: boolean;

  /**
   * Required version
   */
  requiredVersion?: string | false;

  /**
   * Strict version checking
   */
  strictVersion?: boolean;
}

/**
 * Configuration priority:
 * 1. CLI `--mode` parameter (highest priority)
 * 2. `infra.config.js` `mode` field
 * 3. Default: null (standalone mode)
 */
export type InfraConfigPriority = [
  'CLI --mode parameter',
  'infra.config.js mode field',
  'Default null (standalone)'
];

// Export for JSDoc type checking in JavaScript files
export as namespace InfraConfigNS;
