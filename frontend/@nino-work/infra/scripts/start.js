#!/usr/bin/env node

'use strict';

process.env.NODE_ENV = 'development';

const { rsbuild } = require('@rsbuild/core');
const { createRsbuildConfig, loadEnv } = require('../config/rsbuild.config');

loadEnv();

async function start() {
  const config = createRsbuildConfig('development');

  const rsbuildInstance = await rsbuild({
    rsbuildConfig: config,
  });

  const result = await rsbuildInstance.startDevServer();

  const urls = result.urls;
  if (urls && urls.local) {
    console.log(`\n🚀 Dev server running at ${urls.local}\n`);
  }
}

start().catch(err => {
  console.error('Failed to start dev server:', err);
  process.exit(1);
});
