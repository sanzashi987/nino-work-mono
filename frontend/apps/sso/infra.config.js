process.env.PORT = 3001

module.exports = {

  devServer(config) {
    const { proxy } = config

    config.proxy = {
      ...proxy,
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      }
    }
  }
}