module.exports = {
  port: 3002,
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
};
