import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import eslint from '@rollup/plugin-eslint'
import path from 'path'
import { svgBuilder } from './src/plugins/svgBuilder'; 

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    eslint({
      include: 'src/*.+(vue|js|jsx|ts|tsx)',
    }),
    svgBuilder('./src/icons/svg/')
  ],

  resolve: {
    extensions: ['.mjs', '.js', '.jsx', '.json', '.vue'],
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },

  server: {
    host: '127.0.0.1',
    port: 8081,
    strictPort: true,
  },

  build: {
    outDir: 'dist'
  },

  define: {
    'process.env': {
      BASE_API: '/'
    }
  }
})
