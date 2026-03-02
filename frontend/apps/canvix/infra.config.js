const productionInfo = require('./package.json');
const path = require('path');
const prodVersion = productionInfo.version;

module.exports = {
  port: 3003,
  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
  },
  define: {
    __VERSION__: JSON.stringify(prodVersion),
    TARGET_PLATFORM: JSON.stringify(''),
  },
  // mode: 'micro-app',
  devServer(config) {
    const { proxy } = config;

    config.proxy = {
      ...proxy,
      '/backend': {
        target: 'http://localhost:8111',
        changeOrigin: true,
      },
    };
  },
  webpack(config) {
    config.resolve.alias['@dev-coms'] = path.resolve(__dirname, '../../@canvix');
    return config;
  },
};
