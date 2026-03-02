const productionInfo = require('./package.json');
const path = require('path');
const prodVersion = productionInfo.version;

module.exports = {
  port: 3000,
  // mode: 'micro-host',
  devServer(config) {
    const { proxy } = config;

    config.proxy = {
      ...proxy,
      '/backend/storage/v1': {
        target: 'http://localhost:8111',
        changeOrigin: true,
      },
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
    };
  },
};
