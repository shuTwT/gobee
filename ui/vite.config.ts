import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import tailwindcss from '@tailwindcss/vite'
import AutoImport from 'unplugin-auto-import/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import Components from 'unplugin-vue-components/vite'
// import { visualizer } from "rollup-plugin-visualizer";


// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueJsx(),
    vueDevTools(),
    tailwindcss(),
    AutoImport({
      dts:"types/auto-imports.d.ts",
      imports: [
        'vue',
        'vue-router',
        {
          'naive-ui': [
            'useDialog',
            'useMessage',
            'useNotification',
            'useLoadingBar'
          ]
        }
      ]
    }),
    Components({
      dts:"types/components.d.ts",
      types:[{
        from:'vue-router',
        names:['RouterLink','RouterView']
      }],
      resolvers: [NaiveUiResolver()]
    }),
    // visualizer({
    //   open:true
    // })
  ],
  base:'/console/',
  build:{
    minify:true,
    rollupOptions:{
      output:{
        manualChunks:{
          'highlight.js':['highlight.js'],
          'pinia':['pinia'],
          'axios':['axios'],
          'uppy':['@uppy/core','@uppy/dashboard','@uppy/xhr-upload'],
          'vendor':['vue','vue-router','@vueuse/core'],
        }
      }
    }
  },
  server:{
    port: 5379,
    proxy:{
      '/api':{
        target:'http://localhost:13000/api',
        changeOrigin:true,
        secure:false,
        rewrite:(path)=>path.replace(/^\/api/,'')
      }
    }
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
