const commonjs = require('@rollup/plugin-commonjs');
const resolve = require('@rollup/plugin-node-resolve');
const { terser } = require('rollup-plugin-terser');
const replace = require('@rollup/plugin-replace');
const config = {
  input: 'index.js',
  output: {
    file: 'dist/bundle.js',
    format: 'umd',
    name: 'react',
  },
  external: ['react'],
  plugins: [
    resolve(), // Helps Rollup find modules in node_modules
    commonjs(), // Converts CommonJS modules to ES6
    terser({ format: { comments: false } }), // Minify bundle
    replace({
      'process.env.NODE_ENV': JSON.stringify('production'),
      preventAssignment: true,
    }),
    // injectProcessEnv({ NODE_ENV: "production", }),
  ],
};

module.exports = config;
