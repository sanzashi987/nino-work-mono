/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable global-require */

if (process.env.targetPlatform === 'flutter') {
  exports.InteractionService = require('./native').InteractionService;
} else {
  exports.InteractionService = require('./common').InteractionService;
}
