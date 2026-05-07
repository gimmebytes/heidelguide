import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: '.',
  testMatch: '**/*.spec.ts',
  use: {
    baseURL: 'http://localhost:8080',
    screenshot: 'only-on-failure',
  },
  webServer: {
    command: 'cd .. && go run ./cmd/server',
    port: 8080,
    reuseExistingServer: true,
  },
});
