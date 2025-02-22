// process.env.PORT = 3002
process.env.PORT = 3001

module.exports = {
  // also can read from `process.env.NINO_MODE`
  // mode: 'micro-app',
  devServer(config) {
    const { proxy } = config

    config.proxy = {
      ...proxy,
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      }
    }
  },
  // webpack(config) {
  //   config.output.publicPath = 'http://localhost:3000/'
  //   return config
  // }
}