import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: '/ai',
    name: 'AiManagement',
    redirect: '/ai/chat',
    meta: {
      title: 'AI 助手',
      showLink: true,
      roles: ['admin', 'common', 'guest'],
      rank: 15
    },
    children: [
      {
        path: '/ai/chat',
        name: 'AiChat',
        component: () => import('@/views/ai/chat/index.vue'),
        meta: {
          title: 'AI 聊天',
          showLink: true,
          roles: ['admin', 'common', 'guest'],
        },
      },
      {
        path: '/ai/provider',
        name: 'AiProvider',
        component: () => import('@/views/ai/provider/index.vue'),
        meta: {
          title: 'AI 提供商管理',
          showLink: true,
          roles: ['admin', 'common', 'guest'],
        },
      },
    ]
  }
] satisfies RouteRecordRaw[]
