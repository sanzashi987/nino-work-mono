module.exports = {
  port: 3001,
  // mode: 'micro-app',
  devServer(config) {
    const { proxy } = config;

    config.proxy = {
      ...proxy,
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
    };
  },
};
