const productionInfo = require('./package.json');
const prodVersion = productionInfo.version;
process.env.PORT = 3003;
process.env.__VERSION__ = prodVersion;
process.env.TARGET_PLATFORM = '';
module.exports = {
  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
  },

  define: {},
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
  // webpack(config) {
  //   config.output.publicPath = 'http://localhost:3002/'
  //   return config
  // }
};
