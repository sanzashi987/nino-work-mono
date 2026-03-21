/**
 * @type {import('@nino-work/infra/infra.config').InfraConfig}
 */
module.exports = {
  // Development server port
  port: 3000,

  // Running mode: null (standalone), 'micro-host', or 'micro-app'
  // Can be overridden via CLI: pnpm start-mf --mode=micro-host
  mode: 'micro-host',

  // Tailwind CSS configuration
  // Requires: pnpm add -D tailwindcss @tailwindcss/postcss
  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
    theme: {
      extend: {},
    },
  },

  // Global constants defined at build time
  define: {
    __VERSION__: JSON.stringify(require('./package.json').version),
    __DEV__: JSON.stringify(true),
  },

  // Development server configuration
  devServer(config) {
    // Configure proxy for backend API
    config.proxy = {
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
      '/backend/storage': {
        target: 'http://localhost:8111',
        changeOrigin: true,
      },
    };

    // Custom headers
    config.headers = {
      'X-Custom-Header': 'value',
    };
  },

  // Rspack/Webpack configuration override
  webpack(config) {
    // Add path aliases
    config.resolve.alias = {
      ...config.resolve.alias,
      '@': require('path').resolve(__dirname, 'src'),
      '@components': require('path').resolve(__dirname, 'src/components'),
      '@utils': require('path').resolve(__dirname, 'src/utils'),
    };

    // Must return the config
    return config;
  },
};
