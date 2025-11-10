import BaseLayout from '@/layout/index.vue'

export default [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: {},
  },

  {
    path: '/error',
    meta: {},
    children: [
      {
        path: '/error/404',
        name: '404',
        component: () => import('@/views/error/404.vue'),
        meta: {},
      },
      {
        path: '/error/403',
        name: '403',
        component: () => import('@/views/error/403.vue'),
        meta: {},
      },
      {
        path: '/error/500',
        name: '500',
        component: () => import('@/views/error/500.vue'),
        meta: {},
      },
    ],
  },
  {
    path: '/initialize',
    component: () => import('@/views/initialize/index.vue'),
    meta: {},
  },

]
