process.env.PORT = 3003

module.exports = {
  // mode: 'micro-app',
  devServer(config) {
    const { proxy } = config

    config.proxy = {
      ...proxy,
      '/backend': {
        target: 'http://localhost:8111',
        changeOrigin: true,
      }
    }
  },
  // webpack(config) {
  //   config.output.publicPath = 'http://localhost:3002/'
  //   return config
  // }
}