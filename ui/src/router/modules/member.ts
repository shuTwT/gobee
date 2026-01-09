import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: '/member',
    name: 'MemberManagement',
    redirect: '/member/list',
    meta: {
      title: '会员管理',
      showLink: true,
      roles: ['admin'],
      rank: 30
    },
    children:[
      {
        path: '/member/list',
        name: 'MemberList',
        component: () => import('@/views/member/member/index.vue'),
        meta: {
          title: '会员列表',
          showLink: true,
          roles: ['admin'],
          rank: 30
        },
      },
      {
        path: '/member/level',
        name: 'MemberLevelManagement',
        component: () => import('@/views/member/member-level/index.vue'),
        meta: {
          title: '会员等级管理',
          showLink: true,
          roles: ['admin'],
          rank: 30
        },
      }
    ]
  },
] satisfies RouteRecordRaw[]
