import { fileURLToPath, URL } from 'node:url'

import fs from 'fs';
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'

// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vite-plugin
import vuetify from 'vite-plugin-vuetify'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueJsx(),
    vuetify({ autoImport: true })
  ],
  server: {
    host: "127.0.0.1",
    port: 8888,
    https: {
      key: fs.readFileSync('../tmp/ssl/server.key'),
      cert: fs.readFileSync('../tmp/ssl/server.crt')
    }
  },

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
