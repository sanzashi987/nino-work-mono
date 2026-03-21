# @nino-work/infra

基于 Rsbuild 的前端工程化解决方案，为 nino-work monorepo 提供统一的开发和构建能力。

## 核心特性

- ⚡️ **Rsbuild 驱动** - 基于 Rspack，构建速度比 webpack 快 5-10 倍
- ⚛️ **React 生态** - 内置 React 18+、Fast Refresh、React Hook Form
- 📦 **微前端架构** - 支持 Module Federation + Single-spa 模式
- 🎨 **样式方案** - Sass/SCSS、CSS Modules、Tailwind CSS v4
- 📝 **TypeScript** - 开箱即用的 TypeScript 支持和配置文件类型提示
- 🔧 **灵活配置** - 通过 `infra.config.js` 统一配置，带完整类型提示
- 🌍 **环境变量** - 支持 `.env` 文件和 `REACT_APP_*` 变量
- 🖼️ **资源处理** - SVG 转 React 组件、图片/字体优化

## 快速开始

### 安装

在 monorepo 的应用中添加依赖：

```json
{
  "devDependencies": {
    "@nino-work/infra": "workspace:*"
  }
}
```

### npm scripts

```json
{
  "scripts": {
    "start": "infra-start",
    "start-mf": "infra-start --mode=micro-host",
    "build": "infra-build"
  }
}
```

### 目录结构

```
your-app/
├── public/
│   └── index.html          # 可选，自定义 HTML 模板
├── src/
│   ├── index.tsx           # 独立模式入口
│   ├── index.micro.tsx     # 微前端模式入口
│   └── App.tsx             # 应用组件
├── infra.config.js         # 配置文件
├── package.json
└── tsconfig.json
```

## TypeScript 类型提示

`@nino-work/infra` 提供完整的 TypeScript 类型定义，支持在 JavaScript 配置文件中获得类型提示和自动补全。

### 方式 1: 使用 JSDoc 注解（推荐）

在 `infra.config.js` 文件顶部添加 JSDoc 注解：

```javascript
/**
 * @type {import('@nino-work/infra/infra.config').InfraConfig}
 */
module.exports = {
  port: 3000,
  mode: 'micro-host',
  devServer(config) {
    config.proxy = {
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
    };
  },
};
```

现在在编辑器中会获得：
- ✅ 字段名自动补全
- ✅ 类型检查
- ✅ 悬停文档提示
- ✅ 参数类型提示

### 方式 2: 使用 TypeScript 文件

将配置文件改为 `infra.config.ts`：

```typescript
import type { InfraConfig } from '@nino-work/infra/infra.config';

const config: InfraConfig = {
  port: 3000,
  mode: 'micro-host',
  devServer(config) {
    config.proxy = {
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
    };
  },
};

export default config;
```

### 类型定义覆盖

所有配置选项都有完整类型定义：

- `port` - 开发服务器端口
- `mode` - 运行模式（`null | 'micro-host' | 'micro-app'`）
- `tailwind` - Tailwind CSS 配置
- `define` - 编译时常量定义
- `devServer(config)` - 开发服务器配置
- `webpack(config)` - Rspack 配置覆盖

### 示例配置文件

查看 [`infra.config.example.js`](./infra.config.example.js) 获取完整示例。

## 运行模式

支持三种运行模式，通过配置或 CLI 参数指定：

### 1. Standalone 模式（默认）

独立运行的 Web 应用，生成完整的 HTML。

```bash
# 方式 1: 默认行为
pnpm start

# 方式 2: 显式指定（通过 infra.config.js）
module.exports = {
  mode: null  // 或不设置 mode 字段
}
```

**特点：**
- 生成 HTML 文件
- 完整的独立应用
- 适合开发和部署独立页面

**入口文件：** `src/index.tsx`

### 2. Micro-Host 模式

微前端主应用，加载其他子应用。

```bash
# 方式 1: CLI 参数
pnpm start-mf  # --mode=micro-host

# 方式 2: 配置文件
module.exports = {
  mode: 'micro-host'
}
```

**特点：**
- 生成 HTML 文件
- 使用 `scriptLoading: 'module'`
- 通过 single-spa 注册子应用
- 提供 DOM 容器给子应用

**入口文件：** `src/index.micro.tsx`

**典型结构：**
```tsx
// index.micro.tsx
import { registerApplication, start } from 'single-spa';
import { getImportMap } from '@nino-work/mf';

const importMapPromise = getImportMap();

importMapPromise.then(apps => {
  apps.forEach(app => {
    registerApplication(
      app.name,
      () => System.import(app.code),
      location => location.pathname.startsWith(app.path),
      { domElement: () => document.getElementById('sub-app') }
    );
  });
  start();
});

root.render(<App importMapPromise={importMapPromise} />);
```

### 3. Micro-App 模式

微前端子应用，被主应用加载。

```bash
# 方式 1: CLI 参数
pnpm start-mf  # --mode=micro-app

# 方式 2: 配置文件
module.exports = {
  mode: 'micro-app'
}
```

**特点：**
- 不生成 HTML 文件
- 通过 Module Federation 暴露模块
- 被主应用动态加载
- 共享依赖：react、react-dom、single-spa

**入口文件：** `src/index.micro.tsx`

**配置优先级：**
1. CLI `--mode` 参数（最高优先级）
2. `infra.config.js` 中的 `mode` 字段
3. 默认值：`null`（standalone）

## 配置详解：infra.config.js

### 完整配置示例

```javascript
const path = require('path');
const pkg = require('./package.json');

module.exports = {
  // 端口配置
  port: 3000,

  // 运行模式
  mode: 'micro-host',  // null | 'micro-host' | 'micro-app'

  // Tailwind CSS 配置
  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
    theme: {
      extend: {},
    },
  },

  // 全局常量定义
  define: {
    __VERSION__: JSON.stringify(pkg.version),
    __BUILD_TIME__: JSON.stringify(new Date().toISOString()),
    CUSTOM_VAR: JSON.stringify('value'),
  },

  // 开发服务器配置
  devServer(config) {
    config.proxy = {
      '/backend': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
      '/backend/storage': {
        target: 'http://localhost:8111',
        changeOrigin: true,
      },
    };

    // 其他 devServer 配置
    config.headers = {
      'X-Custom-Header': 'value',
    };
  },

  // Rspack/Webpack 配置覆盖
  webpack(config) {
    // 添加别名
    config.resolve.alias = {
      ...config.resolve.alias,
      '@components': path.resolve(__dirname, 'src/components'),
      '@utils': path.resolve(__dirname, 'src/utils'),
    };

    // 自定义插件
    // config.plugins.push(...)

    return config;  // 必须返回 config
  },
};
```

### 配置选项说明

#### port
- 类型: `number`
- 默认: `3000`
- 说明: 开发服务器端口号

#### mode
- 类型: `null | 'micro-host' | 'micro-app'`
- 默认: `null`
- 说明: 运行模式，可被 CLI `--mode` 参数覆盖

#### tailwind
- 类型: `object`
- 默认: `undefined`
- 说明: Tailwind CSS 配置，传递给 `@tailwindcss/postcss`
- 要求: 需要安装 `tailwindcss` 和 `@tailwindcss/postcss`

#### define
- 类型: `object`
- 默认: `{}`
- 说明: 编译时全局常量定义
- 示例: `{ __DEV__: JSON.stringify(true) }`

#### devServer(config)
- 类型: `(config: DevServerConfig) => void`
- 说明: 自定义开发服务器配置
- 常用: 配置 proxy、headers、cors 等
- 注意: 直接修改 config 对象，无需返回

#### webpack(config)
- 类型: `(config: RspackConfig) => RspackConfig | undefined`
- 说明: 自定义 Rspack 配置
- 用途: 添加别名、插件、修改输出等
- 注意: 必须返回修改后的 config

## 入口文件模式

### 双入口文件

每个应用通常有两个入口文件：

#### 1. index.tsx（独立模式）

```tsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import './index.scss';

// 独立运行，不加载微前端
const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(<App importMapPromise={Promise.resolve([])} />);
```

#### 2. index.micro.tsx（微前端模式）

```tsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import { registerApplication, start } from 'single-spa';
import { getImportMap } from '@nino-work/mf';
import App from './App';
import './index.scss';

// 加载微前端配置
const importMapPromise = getImportMap();

// 注册子应用（仅 micro-host）
importMapPromise.then(apps => {
  apps.forEach(app => {
    registerApplication(
      app.name,
      () => System.import(/* webpackIgnore: true */ app.code),
      location => location.pathname.startsWith(app.path),
      () => ({
        domElement: () => document.getElementById('nino-sub-app'),
        basename: app.path.startsWith('/') ? app.path.slice(1) : app.path,
      })
    );
  });
  start();
});

// 渲染应用
const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(<App importMapPromise={importMapPromise} />);
```

### App.tsx 共享组件

```tsx
import React from 'react';
import { usePromise } from '@nino-work/shared';
import { MenuMeta } from '@nino-work/mf';

type AppProps = {
  importMapPromise: Promise<MenuMeta[]>;
};

const App: React.FC<AppProps> = ({ importMapPromise }) => {
  const { data } = usePromise(importMapPromise);
  const isLoading = data === null;

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <ThemeProvider>
      <RouterProvider router={router} />
    </ThemeProvider>
  );
};

export default App;
```

## 环境变量

### .env 文件支持

支持以下 `.env` 文件，按优先级加载：

1. `.env.[NODE_ENV].local` - 环境特定本地配置（最高优先级）
2. `.env.local` - 本地配置
3. `.env.[NODE_ENV]` - 环境特定配置
4. `.env` - 默认配置

**示例：**
```bash
# .env.development
REACT_APP_API_URL=http://localhost:8081
REACT_APP_ENV=development

# .env.production
REACT_APP_API_URL=https://api.example.com
REACT_APP_ENV=production
```

### 内置环境变量

- `NODE_ENV` - `development` / `production` / `test`
- `PUBLIC_URL` - 从 `package.json` 的 `homepage` 字段读取

### 自定义环境变量

所有 `REACT_APP_*` 开头的变量会自动注入：

```javascript
// 在代码中使用
console.log(process.env.REACT_APP_API_URL);
console.log(process.env.NODE_ENV);
```

### 编译时常量

通过 `define` 配置定义：

```javascript
// infra.config.js
module.exports = {
  define: {
    __VERSION__: JSON.stringify(require('./package.json').version),
  }
}

// 代码中使用
console.log(__VERSION__);  // 编译时替换为字符串字面量
```

## HTML 模板

### 默认模板

位于 `@nino-work/infra/config/template.html`：

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="theme-color" content="#000000" />
    <meta name="description" content="<%= title %>" />
    <title><%= title %></title>
  </head>
  <body>
    <noscript>You need to enable JavaScript to run this app.</noscript>
    <div id="root"></div>
  </body>
</html>
```

### 模板变量

- `<%= title %>` - 应用标题（从 `package.json` name 字段读取）
- `<%= description %>` - 描述（从 `package.json` description 字段读取）

### 自定义模板

在应用根目录创建 `public/index.html`：

```html
<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8" />
    <title>自定义标题</title>
  </head>
  <body>
    <div id="root"></div>
    <div id="nino-sub-app"></div>  <!-- 微前端容器 -->
  </body>
</html>
```

## Module Federation 配置

### 子应用配置（micro-app）

```javascript
// infra.config.js
module.exports = {
  mode: 'micro-app',
  // 自动配置 Module Federation
  // name: package.json 的 name（非字母数字替换为 _）
  // shared: react, react-dom, single-spa
}
```

**生成配置：**
```javascript
moduleFederationPlugin({
  name: 'my_app',  // 从 package.json name 转换
  exposes: {},
  shared: {
    react: { singleton: true, eager: true },
    'react-dom': { singleton: true, eager: true },
    'single-spa': { singleton: true },
  },
})
```

### 主应用配置（micro-host）

主应用不需要特殊的 Module Federation 配置，只需：

```javascript
// infra.config.js
module.exports = {
  mode: 'micro-host',
  devServer(config) {
    config.proxy = {
      // 代理到后端服务
    };
  }
}
```

### System.js 使用

在主应用中动态加载子应用：

```javascript
// 动态导入 Module Federation 模块
const module = await System.import(/* webpackIgnore: true */ 'canvix_app');
```

## 样式处理

### Sass/SCSS

内置支持，无需配置：

```tsx
import './styles.scss';  // 直接导入
```

### CSS Modules

文件命名 `[name].module.css`：

```tsx
import styles from './Button.module.css';

<button className={styles.button}>Click</button>
```

### Tailwind CSS

#### 安装依赖

```bash
pnpm add -D tailwindcss @tailwindcss/postcss
```

#### 配置

```javascript
// infra.config.js
module.exports = {
  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
  }
}
```

#### 使用

```tsx
// index.tsx 或 index.css
import 'tailwindcss';
```

## 资源处理

### 图片

```tsx
import logo from './logo.png';  // string: URL

<img src={logo} alt="Logo" />
```

### SVG 转 React 组件

```tsx
import { ReactComponent as Logo } from './logo.svg';

<Logo width={100} height={100} />
```

或：

```tsx
import Logo from './logo.svg?react';  // 返回 React 组件

<Logo />
```

### 字体

```css
@font-face {
  font-family: 'CustomFont';
  src: url('./fonts/custom.woff2') format('woff2');
}
```

## TypeScript 支持

### tsconfig.json

```json
{
  "extends": "@nino-work/infra/tsconfig.base.json",
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"]
    }
  },
  "include": ["src"]
}
```

### 类型声明

内置类型声明文件 `react-app-env.d.ts`：

```typescript
/// <reference types="@rsbuild/core" />

declare namespace NodeJS {
  interface ProcessEnv {
    NODE_ENV: 'development' | 'production' | 'test';
    PUBLIC_URL: string;
    REACT_APP_*: string;
  }
}
```

## 构建输出

### 目录结构

```
build/
├── static/
│   ├── js/
│   │   ├── main.[contenthash:8].js
│   │   └── vendors.[contenthash:8].js
│   ├── css/
│   │   └── main.[contenthash:8].css
│   └── media/
│       ├── logo.[hash:8].png
│       └── font.[hash:8].woff2
└── index.html  # standalone 和 micro-host 模式
```

### 文件命名

- **开发模式**: `[name].js` / `[name].css`
- **生产模式**: `[name].[contenthash:8].js` / `[name].[contenthash:8].css`

### publicPath

从 `package.json` 的 `homepage` 字段读取：

```json
{
  "homepage": "/app/"
}
```

构建后所有资源路径带 `/app/` 前缀。

## 常见问题

### 1. 端口冲突

**问题：** 默认端口 3000 已被占用

**解决：**
```javascript
// infra.config.js
module.exports = {
  port: 3001
}
```

### 2. 代理不生效

**问题：** 后端 API 请求 404

**解决：**
```javascript
module.exports = {
  devServer(config) {
    config.proxy = {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: { '^/api': '' },  // 可选：重写路径
      },
    };
  }
}
```

### 3. 微前端模式无法切换

**问题：** `pnpm start-mf` 仍然运行独立模式

**解决：**
```javascript
// 确保 infra.config.js 导出 mode
module.exports = {
  mode: 'micro-host'  // 或 'micro-app'
}

// 或者通过 CLI
// "start-mf": "infra-start --mode=micro-host"
```

### 4. Tailwind 样式不生效

**问题：** Tailwind 类名不生效

**解决：**
```bash
# 1. 安装依赖
pnpm add -D tailwindcss @tailwindcss/postcss

# 2. 配置 infra.config.js
module.exports = {
  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
  }
}

# 3. 导入 Tailwind
// index.tsx 或 index.css
import 'tailwindcss';
```

### 5. SVG 导入报错

**问题：** `import Logo from './logo.svg'` 报错

**解决：**
```tsx
// 方式 1: 作为 URL
import logoUrl from './logo.svg';

// 方式 2: 作为 React 组件
import { ReactComponent as Logo } from './logo.svg';

// 方式 3: 使用 ?react 后缀
import Logo from './logo.svg?react';
```

## 最佳实践

### 1. 配置文件组织

```javascript
// infra.config.js
const path = require('path');
const pkg = require('./package.json');

const isProduction = process.env.NODE_ENV === 'production';

module.exports = {
  port: process.env.PORT || 3000,

  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
  },

  define: {
    __VERSION__: JSON.stringify(pkg.version),
    __DEV__: JSON.stringify(!isProduction),
  },

  devServer(config) {
    if (!isProduction) {
      config.proxy = {
        '/backend': {
          target: 'http://localhost:8081',
          changeOrigin: true,
        },
      };
    }
  },

  webpack(config) {
    config.resolve.alias['@'] = path.resolve(__dirname, 'src');
    return config;
  },
};
```

### 2. 环境变量管理

```bash
# .env.development.local (本地开发，不提交 git)
REACT_APP_API_URL=http://localhost:8081
REACT_APP_DEBUG=true

# .env.production (提交 git)
REACT_APP_API_URL=https://api.example.com
REACT_APP_DEBUG=false
```

### 3. 代码分割

Rsbuild 自动进行代码分割，也可手动：

```tsx
import { lazy, Suspense } from 'react';

const Dashboard = lazy(() => import('./pages/Dashboard'));

<Suspense fallback={<div>Loading...</div>}>
  <Dashboard />
</Suspense>
```

### 4. 性能优化

```javascript
// infra.config.js
module.exports = {
  webpack(config) {
    // 生产模式优化
    if (process.env.NODE_ENV === 'production') {
      config.optimization = {
        ...config.optimization,
        splitChunks: {
          chunks: 'all',
          cacheGroups: {
            vendors: {
              test: /[\\/]node_modules[\\/]/,
              name: 'vendors',
            },
          },
        },
      };
    }
    return config;
  }
}
```

## 迁移指南

### 从 webpack 版本迁移

v0.2.0 从 webpack 迁移到 Rsbuild，主要变化：

#### 1. 依赖精简

移除以下依赖（已内置或不再需要）：

```json
{
  "devDependencies": {
    // 不再需要
    "@babel/core": "...",
    "@babel/preset-env": "...",
    "@babel/preset-react": "...",
    "@babel/preset-typescript": "...",
    "babel-loader": "...",
    "css-loader": "...",
    "sass-loader": "...",
    "style-loader": "...",
    "file-loader": "...",
    "url-loader": "...",
    "react-refresh": "...",
    "@pmmmwh/react-refresh-webpack-plugin": "...",
    "html-webpack-plugin": "...",
    "mini-css-extract-plugin": "...",
    "terser-webpack-plugin": "...",
    "webpack": "...",
    "webpack-cli": "...",
    "webpack-dev-server": "..."
  }
}
```

#### 2. 配置简化

**旧配置（webpack）：**
```javascript
// 复杂的 webpack 配置
module.exports = {
  module: {
    rules: [
      { test: /\.tsx?$/, use: 'babel-loader' },
      { test: /\.css$/, use: ['style-loader', 'css-loader'] },
      { test: /\.scss$/, use: ['style-loader', 'css-loader', 'sass-loader'] },
      { test: /\.svg$/, use: ['@svgr/webpack'] },
      // ...更多 loader
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({ template: '...' }),
    new MiniCssExtractPlugin({ ... }),
    // ...更多插件
  ]
}
```

**新配置（Rsbuild）：**
```javascript
// 简洁的 infra.config.js
module.exports = {
  port: 3000,
  // Rsbuild 内置大部分配置
}
```

#### 3. API 兼容

`infra.config.js` 的 API 保持兼容：

- ✅ `port` - 端口配置
- ✅ `mode` - 运行模式
- ✅ `define` - 全局常量
- ✅ `devServer` - 开发服务器配置
- ✅ `webpack` - Rspack 配置覆盖（API 兼容 webpack）
- ✅ `tailwind` - Tailwind 配置

#### 4. 性能提升

- 构建速度提升 5-10 倍
- 热更新速度显著提升
- 内存占用减少

## 技术栈

### 核心依赖

- `@rsbuild/core` - 构建核心
- `@rsbuild/plugin-react` - React 支持
- `@rsbuild/plugin-sass` - Sass/SCSS 支持
- `@rsbuild/plugin-svgr` - SVG 转 React 组件
- `@module-federation/rsbuild-plugin` - Module Federation 支持

### 开发依赖

- `dotenv` - 环境变量加载
- `dotenv-expand` - 环境变量扩展
- `fs-extra` - 文件系统工具

### Peer Dependencies

- `tailwindcss` >= 4.0.0 - Tailwind CSS（可选）
- `@tailwindcss/postcss` >= 4.0.0 - Tailwind PostCSS 插件（可选）

## License

MIT

---

**文档版本**: v0.2.0
**最后更新**: 2026-03-21
**维护者**: nino-work team
