import { fileURLToPath, URL } from 'node:url'
import { resolve, dirname } from 'node:path'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import VueI18nPlugin from '@intlify/unplugin-vue-i18n/vite'
// import commonjs from 'vite-plugin-commonjs'
// import commonjs from "vite-plugin-commonjs";
// import { nodePolyfills } from 'vite-plugin-node-polyfills';

// https://vite.dev/config/
export default defineConfig({
  define: {
    'process.env': {}
  },
  plugins: [
    // commonjs(),
    // nodePolyfills({
    //   // Include all Node.js polyfills
    //   include: ['process'],
    // }),
    // commonjs(),
    vue(),
    vueDevTools(),
    VueI18nPlugin({
      runtimeOnly: false,
      // include: resolve(dirname(fileURLToPath(import.meta.url)), './src/i18n/locales/**'),
    })
  ],

  // build: {
  //   commonjsOptions: {
  //     transformMixedEsModules: true
  //   },
  // },
  //   rollupOptions: {
  //     external: ['nock', 'aws-sdk', 'mock-aws-s3']
  //   },
  //   commonjsOptions: {
  //     transformMixedEsModules: true
  //   },
  // },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
