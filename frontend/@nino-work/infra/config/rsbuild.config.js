const fs = require('fs');
const path = require('path');
const { defineConfig } = require('@rsbuild/core');
const pluginReact = require('@rsbuild/plugin-react');
const pluginSass = require('@rsbuild/plugin-sass');
const pluginSvgr = require('@rsbuild/plugin-svgr');
const { moduleFederationPlugin } = require('@module-federation/rsbuild-plugin');

const appDirectory = fs.realpathSync(process.cwd());
const resolveApp = relativePath => path.resolve(appDirectory, relativePath);

// Default HTML template in infra package
const defaultHtmlTemplate = path.resolve(__dirname, 'template.html');

const paths = {
  dotenv: resolveApp('.env'),
  appPath: resolveApp('.'),
  appBuild: resolveApp(process.env.BUILD_PATH || 'build'),
  appPublic: resolveApp('public'),
  appHtml: fs.existsSync(resolveApp('public/index.html')) ? resolveApp('public/index.html') : defaultHtmlTemplate,
  appIndexJs: resolveModule(resolveApp, 'src/index'),
  appIndexMicro: resolveModule(resolveApp, 'src/index.micro'),
  appPackageJson: resolveApp('package.json'),
  appSrc: resolveApp('src'),
  appTsConfig: resolveApp('tsconfig.json'),
  publicUrlOrPath: getPublicUrlOrPath(),
  infraPath: resolveApp('infra.config.js'),
};

function getPublicUrlOrPath() {
  const pkg = require(resolveApp('package.json'));
  return pkg.homepage || '/';
}

const moduleFileExtensions = ['tsx', 'ts', 'jsx', 'js', 'json'];

function resolveModule(resolveFn, filePath) {
  const extension = moduleFileExtensions.find(ext =>
    fs.existsSync(resolveFn(`${filePath}.${ext}`))
  );
  return extension ? resolveFn(`${filePath}.${extension}`) : resolveFn(`${filePath}.js`);
}

function loadEnv() {
  const NODE_ENV = process.env.NODE_ENV || 'development';
  const dotenvFiles = [
    `${paths.dotenv}.${NODE_ENV}.local`,
    NODE_ENV !== 'test' && `${paths.dotenv}.local`,
    `${paths.dotenv}.${NODE_ENV}`,
    paths.dotenv,
  ].filter(Boolean);

  dotenvFiles.forEach(dotenvFile => {
    if (fs.existsSync(dotenvFile)) {
      require('dotenv-expand').expand(
        require('dotenv').config({ path: dotenvFile })
      );
    }
  });
}

function getEnvDefinitions() {
  const REACT_APP = /^REACT_APP_/i;
  const raw = Object.keys(process.env)
    .filter(key => REACT_APP.test(key))
    .reduce(
      (env, key) => {
        env[key] = process.env[key];
        return env;
      },
      {
        NODE_ENV: process.env.NODE_ENV || 'development',
        PUBLIC_URL: paths.publicUrlOrPath,
      }
    );
  return raw;
}

function loadInfraConfig() {
  if (fs.existsSync(paths.infraPath)) {
    return require(paths.infraPath);
  }
  return {};
}

function createRsbuildConfig(webpackEnv) {
  loadEnv();

  const isEnvDevelopment = webpackEnv === 'development';
  const isEnvProduction = webpackEnv === 'production';
  const env = getEnvDefinitions();
  const infraConfig = loadInfraConfig();

  const isMicro = infraConfig.mode === 'micro-app' || infraConfig.mode === 'micro-host';
  const isMicroHost = infraConfig.mode === 'micro-host';

  const pkg = require(paths.appPackageJson);

  const baseConfig = {
    source: {
      entry: isMicro ? paths.appIndexMicro : paths.appIndexJs,
      define: {
        ...infraConfig.define,
        ...Object.keys(env).reduce((acc, key) => {
          acc[`process.env.${key}`] = JSON.stringify(env[key]);
          return acc;
        }, {}),
        NINO_IS_PROD: isEnvProduction,
      },
    },
    output: {
      distPath: {
        root: paths.appBuild,
        js: 'static/js',
        css: 'static/css',
        svg: 'static/media',
        font: 'static/media',
        image: 'static/media',
        media: 'static/media',
      },
      filename: {
        js: isEnvProduction ? '[name].[contenthash:8].js' : '[name].js',
        css: isEnvProduction ? '[name].[contenthash:8].css' : '[name].css',
      },
      assetPrefix: paths.publicUrlOrPath,
    },
    html: !isMicro || isMicroHost ? {
      template: paths.appHtml,
      scriptLoading: isMicroHost ? 'module' : 'defer',
    } : false,
    devServer: isEnvDevelopment ? {
      port: parseInt(process.env.PORT, 10) || 3000,
      host: '0.0.0.0',
      hot: true,
      historyApiFallback: true,
    } : undefined,
    tools: {
      rspack: {
        resolve: {
          extensions: ['.tsx', '.ts', '.jsx', '.js', '.json'],
          alias: {
            'react-native': 'react-native-web',
          },
        },
      },
    },
    plugins: [
      pluginReact(),
      pluginSass(),
      pluginSvgr({
        svgrOptions: {
          prettier: false,
          svgo: false,
          titleProp: true,
          ref: true,
        },
      }),
    ],
  };

  // Module Federation for micro-app mode
  if (isMicro && !isMicroHost) {
    baseConfig.plugins.push(
      moduleFederationPlugin({
        name: pkg.name.replace(/[^a-zA-Z0-9]/g, '_'),
        exposes: {},
        shared: {
          react: { singleton: true, eager: true },
          'react-dom': { singleton: true, eager: true },
          'single-spa': { singleton: true },
        },
      })
    );
  }

  // Apply infra config overrides
  let finalConfig = baseConfig;

  // devServer override
  if (infraConfig.devServer && isEnvDevelopment) {
    const serverConfig = { ...finalConfig.devServer };
    infraConfig.devServer(serverConfig);
    finalConfig.devServer = serverConfig;
  }

  // rspack/webpack override
  if (infraConfig.webpack) {
    const rspackConfig = { ...finalConfig.tools.rspack };
    const overridden = infraConfig.webpack({ ...rspackConfig, output: { ...rspackConfig.output } });
    if (overridden) {
      finalConfig.tools.rspack = overridden;
    }
  }

  // Tailwind support via PostCSS
  if (infraConfig.tailwind) {
    const tailwindConfig = infraConfig.tailwind;
    finalConfig.tools = finalConfig.tools || {};
    finalConfig.tools.postcss = {
      postcssOptions: {
        plugins: [
          ['@tailwindcss/postcss', tailwindConfig],
        ],
      },
    };
  }

  return defineConfig(finalConfig);
}

module.exports = {
  createRsbuildConfig,
  paths,
  loadEnv,
  getEnvDefinitions,
  loadInfraConfig,
};
