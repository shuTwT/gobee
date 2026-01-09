import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: '/ai',
    name: 'AiManagement',
    redirect: '/ai/chat',
    meta: {
      title: 'AI 助手',
      showLink: true,
      roles: ['admin', 'common'],
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
          roles: ['admin', 'common'],
        },
      },
    ]
  }
] satisfies RouteRecordRaw[]
