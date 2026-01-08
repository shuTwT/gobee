import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: '/member',
    name: 'MemberManagement',
    component: () => import('@/views/member/index.vue'),
    meta: {
      title: '会员管理',
      showLink: true,
      roles: ['admin'],
      rank: 30
    },
  },
] satisfies RouteRecordRaw[]
