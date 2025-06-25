import { defineConfig, loadEnv } from 'vite'
import vuePlugin from '@vitejs/plugin-vue'
import path from 'path'
import { svgBuilder } from './src/components/SvgBuilder'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import ElementPlus from 'unplugin-element-plus/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

// config detail in https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  process.env.NODE_ENV = env.VITE_NODE_ENV || 'dev'
  return {
    plugins: [
      vuePlugin(),
      svgBuilder('./src/components/icons/svg/'),
      ElementPlus({}),
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
    ],

    resolve: {
      extensions: ['.mjs', '.js', '.jsx', '.json', '.vue'],
      alias: {
        '~': path.resolve(__dirname, './src'),
      },
    },

    server: {
      host: '0.0.0.0',
      port: 8088,
      strictPort: true,
      proxy: {
        '/api': {
          target: env.VITE_BASE_API || 'http://localhost:3000',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '')
        }
      }
    },

    build: {
      outDir: 'dist',
      chunkSizeWarningLimit: 768,
      rollupOptions: {
        output: {
          manualChunks(id) {
            if (id.includes('node_modules')) {
              return id.toString().split('node_modules/')[1].split('/')[0].toString()
            }
          }
        }
      }
    },

    define: {
      'process.env': {
        NODE_ENV: env.VITE_NODE_ENV,
        BASE_API: env.VITE_BASE_API,
      }
    },

    css: {
      preprocessorOptions: {
        scss: {
          silenceDeprecations: ['legacy-js-api'], // remove when sass >= 2.0.0
        }
      }
    }
  }
})
