import { defineConfig, loadEnv } from 'vite'
import vuePlugin from '@vitejs/plugin-vue'
import path from 'path'
import { svgBuilder } from './src/plugins/svgBuilder'

// config detail in https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  process.env.NODE_ENV = env.VITE_NODE_ENV || 'dev'
  return {
    plugins: [
      vuePlugin(),
      svgBuilder('./src/icons/svg/')
    ],

    resolve: {
      extensions: ['.mjs', '.js', '.jsx', '.json', '.vue'],
      alias: {
        '@': path.resolve(__dirname, './src'),
      },
    },

    server: {
      host: '0.0.0.0',
      port: 8088,
      strictPort: true,
    },

    build: {
      outDir: 'dist'
    },

    define: {
      'process.env': {
        BASE_API: env.VITE_BASE_API,
      }
    },
  }
})
