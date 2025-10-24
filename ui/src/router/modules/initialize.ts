import type { RouteRecordRaw } from 'vue-router';
import Layout from '@/layout/index.vue';

const initializeRouter: RouteRecordRaw[] = [
  {
    path: '/initialize',
    component: Layout,
    redirect: '/initialize/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/initialize/index.vue'),
        name: 'Initialize',
        meta: {
          title: '系统初始化',
          hidden: true
        }
      }
    ]
  }
];

export default initializeRouter;