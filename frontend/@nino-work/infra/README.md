# @nino-work/infra

基于 Rsbuild 的构建工具封装，为 nino-work monorepo 提供统一的开发构建能力。

## 特性

- ⚡️ **Rsbuild 驱动** - 基于 Rspack，构建速度提升 5-10 倍
- ⚛️ **React 支持** - 内置 React 18+ 和 Fast Refresh
- 📦 **微前端支持** - 支持 Module Federation 模式
- 🎨 **样式支持** - Sass/SCSS、CSS Modules、Tailwind CSS
- 📝 **TypeScript** - 开箱即用的 TypeScript 支持
- 🔧 **灵活配置** - 通过 `infra.config.js` 自定义

## 使用方法

### 安装

在 monorepo 的 app 中添加依赖：

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
    "build": "infra-build"
  }
}
```

### 目录结构

```
your-app/
├── public/
│   └── index.html      # 可选，自定义 HTML 模板
├── src/
│   ├── index.tsx       # 入口文件
│   └── index.micro.tsx # 微前端入口（可选）
├── infra.config.js     # 配置文件
└── package.json
```

## 配置文件 (infra.config.js)

### 基础配置

```javascript
module.exports = {
  // 开发服务器端口
  // process.env.PORT = 3000;
  
  // Tailwind CSS 配置
  tailwind: {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
  },
  
  // 环境变量定义
  define: {
    MY_VAR: JSON.stringify('value'),
  },
};
```

### 微前端模式

```javascript
// 主应用 (Host)
module.exports = {
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

// 子应用 (Remote)
module.exports = {
  mode: 'micro-app',
};
```

### 自定义 devServer

```javascript
module.exports = {
  devServer(config) {
    config.proxy = {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    };
  },
};
```

### 自定义 Rspack 配置

```javascript
module.exports = {
  webpack(config) {
    config.resolve.alias['@components'] = path.resolve(__dirname, 'src/components');
    return config;
  },
};
```

## 迁移自 webpack 版本

v0.2.0 从 webpack 迁移到 Rsbuild，主要变化：

1. **依赖精简** - 移除了大量 babel/loader 相关依赖
2. **配置简化** - Rsbuild 内置大部分常用配置
3. **性能提升** - 构建速度提升 5-10 倍
4. **兼容性** - `infra.config.js` API 保持兼容

### 不再需要的依赖

以下依赖已内置或不再需要：
- `@babel/*` 系列
- `*-loader` 系列（css-loader, sass-loader 等）
- `react-refresh` / `@pmmmwh/react-refresh-webpack-plugin`
- `html-webpack-plugin`
- `mini-css-extract-plugin`
- `terser-webpack-plugin`
- 等等...

## License

MIT
