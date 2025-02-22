process.env.PORT = 3000

module.exports = {
  // mode: 'micro-host',
  devServer(config) {
    const { proxy } = config

    config.proxy = {
      ...proxy,
      '/backend/storage/v1': {
        target: 'http://localhost:8111',
        changeOrigin: true,
      },
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      }
    }
  }
}