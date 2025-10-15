import { createRouter, createWebHistory } from 'vue-router'
import BaseLayout from '../layout/index.vue'
console.log(import.meta.env.BASE_URL)
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/login/index.vue'),
    },
    {
      path: '/',
      name: 'System',
      component: BaseLayout,
      children: [
        {
          path: '/',
          name: 'Home',
          component: () => import('../views/home/index.vue'),
        },
        {
          path: '/album',
          name: 'AlbumManagement',
          component: () => import('../views/album/index.vue'),
        },
        {
          path: '/comment',
          name: 'CommentManagement',
          component: () => import('../views/comment/index.vue'),
        },
        {
          path: '/file',
          name: 'FileManagement',
          component: () => import('../views/file/index.vue'),
        },
        {
          path: '/flink',
          name: 'FlinkManagement',
          component: () => import('../views/flink/index.vue'),
        },
        {
          path: '/model-content',
          name: 'ModelContentManagement',
          component: () => import('../views/model-content/index.vue'),
        },
        {
          path: '/model-schema',
          name: 'ModelSchemaManagement',
          component: () => import('../views/model-schema/index.vue'),
        },
        {
          path: '/pay-channel',
          name: 'PayChannelManagement',
          component: () => import('../views/pay-channel/index.vue'),
        },
        {
          path: '/pay-order',
          name: 'PayOrderManagement',
          component: () => import('../views/pay-order/index.vue'),
        },
        {
          path: '/user',
          name: 'UserManagement',
          component: () => import('../views/user/index.vue'),
        },
        {
          path: '/webhook',
          name: 'WebhookManagement',
          component: () => import('../views/webhook/index.vue'),
        },
      ],
    },

    {
      path: '/:pathMatch(.*)*',
      component: BaseLayout,
      children: [
        {
          path: '',
          name: 'NotFound',
          component: () => import('../views/error/404.vue'),
        },
      ],
    },
    {
      path: '/403',
      component: BaseLayout,
      children: [
        {
          path: '',
          name: 'Forbidden',
          component: () => import('../views/error/403.vue'),
        },
      ],
    },
  ],
})

export default router
