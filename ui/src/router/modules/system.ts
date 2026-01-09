import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: '/system',
    name: 'SystemManagement',
    redirect: '/system/settings',
    meta: {
      title: '系统管理',
      showLink: true,
      roles: ['admin'],
      rank:98
    },
    children: [
      {
        path: '/system/settings',
        name: 'SystemSettings',
        component: () => import('@/views/system/settings/index.vue'),
        meta: {
          title: '系统设置',
          showLink: true,
          roles: ['admin'],
        },
      },
      {
        path: '/system/api',
        name: 'ApiManagement',
        component: () => import('@/views/system/api/index.vue'),
        meta: {
          title: 'api管理',
          showLink: true,
          roles: ['admin'],
        },
      },
      {
        path: '/system/webhook',
        name: 'WebhookManagement',
        component: () => import('@/views/webhook/index.vue'),
        meta: {
          title: 'webhook管理',
          showLink: true,
          roles: ['admin'],
        },
      },
      {
        path: '/system/role',
        name: 'RoleManagement',
        component: () => import('@/views/system/role/index.vue'),
        meta: {
          title: '角色管理',
          showLink: true,
          roles: ['admin'],
        },
      },
      {
        path: '/system/user',
        name: 'UserManagement',
        component: () => import('@/views/system/user/index.vue'),
        meta: {
          title: '用户管理',
          showLink: true,
          roles: ['admin'],
        },
      },
    ],
  },
] satisfies RouteRecordRaw[]
