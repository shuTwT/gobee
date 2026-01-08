import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: '/product',
    name: 'ProductManagement',
    component: () => import('@/views/product/index.vue'),
    meta: {
      title: '商品管理',
      showLink: true,
      roles: ['admin'],
    },
  },
] satisfies RouteRecordRaw[]
