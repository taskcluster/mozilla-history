import { defineConfig } from '@playwright/test'

export default defineConfig({
  testDir: '.',
  testMatch: 'layout.spec.mjs',
  outputDir: 'test-results',
  reporter: 'list',
  use: {
    baseURL: 'http://127.0.0.1:8766',
    browserName: 'firefox',
    screenshot: 'only-on-failure',
    trace: 'retain-on-failure',
  },
  webServer: {
    command: 'python3 -m http.server 8766 --bind 127.0.0.1 --directory ..',
    url: 'http://127.0.0.1:8766/docs/index.html?local',
    reuseExistingServer: false,
  },
})
