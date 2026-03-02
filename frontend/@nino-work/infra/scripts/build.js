#!/usr/bin/env node

'use strict';

process.env.NODE_ENV = 'production';

const { rsbuild } = require('@rsbuild/core');
const { createRsbuildConfig, loadEnv } = require('../config/rsbuild.config');
const fs = require('fs-extra');
const path = require('path');

loadEnv();

async function build() {
  const config = createRsbuildConfig('production');
  const buildPath = config.output?.distPath?.root || 'build';

  // Clean build directory
  await fs.emptyDir(buildPath);

  console.log('Creating an optimized production build...\n');

  const rsbuildInstance = await rsbuild({
    rsbuildConfig: config,
  });

  const result = await rsbuildInstance.build();

  if (result && result.close) {
    await result.close();
  }

  console.log('\n✅ Build completed successfully!\n');
  console.log(`Build output: ${path.resolve(buildPath)}\n`);
}

build().catch(err => {
  console.error('Build failed:', err);
  process.exit(1);
});
