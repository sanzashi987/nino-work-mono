process.env.PORT = 3002

module.exports = {

  devServer(config) {
    const { proxy } = config

    config.proxy = {
      ...proxy,
      '/backend': {
        target: 'http://localhost:8111',
        changeOrigin: true,
      }
    }
  }
}