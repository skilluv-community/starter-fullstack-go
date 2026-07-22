import { defineConfig } from '@playwright/test';

export default defineConfig({
  webServer: {
    command: 'npm run build && node build/index.js',
    port: 3000,
    reuseExistingServer: !process.env.CI,
    timeout: 120_000
  },
  testDir: 'tests/e2e',
  use: { baseURL: 'http://localhost:3000' }
});
