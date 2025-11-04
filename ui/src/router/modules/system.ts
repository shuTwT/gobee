import type { RouteRecordRaw } from 'vue-router'

export default [
    {
      path:"/post",
      name:"PostManagement",
      component:()=>import('@/views/post/index.vue'),
      meta:{
        title:'文章管理',
        showLink:true,
        roles:['admin','common']
      }
    },
    {
      path:"/post/edit",
      name:"PostEditor",
      component:()=>import('@/views/post/editor/index.vue'),
      meta:{
        title:"文章编辑",
        showLink:false,
        roles:['admin','common']
      }
    },
    {
      path: '/album',
      name: 'AlbumManagement',
      component: () => import('@/views/album/index.vue'),
      meta: {
        title: '相册管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: '/comment',
      name: 'CommentManagement',
      component: () => import('@/views/comment/index.vue'),
      meta: {
        title: '评论管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: "/storage-strategy",
      name:"StorageStrategyManagement",
      component:()=>import('@/views/storage-strategy/index.vue'),
      meta:{
        title:'存储策略管理',
        showLink:true,
        roles:['admin','common']
      }
    },
    {
      path: '/file',
      name: 'FileManagement',
      component: () => import('@/views/file/index.vue'),
      meta: {
        title: '文件管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: '/flink',
      name: 'FlinkManagement',
      component: () => import('@/views/flink/index.vue'),
      meta: {
        title: '友链管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    // {
    //   path: '/model-content',
    //   name: 'ModelContentManagement',
    //   component: () => import('@/views/model-content/index.vue'),
    //   meta: {
    //     title: '模型内容管理',
    //     showLink: true,
    //     roles: ['admin', 'common'],
    //   },
    // },
    // {
    //   path: '/model-schema',
    //   name: 'ModelSchemaManagement',
    //   component: () => import('@/views/model-schema/index.vue'),
    //   meta: {
    //     title: '模型内容管理',
    //     showLink: true,
    //     roles: ['admin', 'common'],
    //   },
    // },
    {
      path: '/pay-channel',
      name: 'PayChannelManagement',
      component: () => import('@/views/pay-channel/index.vue'),
      meta: {
        title: '支付渠道管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: '/pay-order',
      name: 'PayOrderManagement',
      component: () => import('@/views/pay-order/index.vue'),
      meta: {
        title: '支付订单管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: '/user',
      name: 'UserManagement',
      component: () => import('@/views/system/user/index.vue'),
      meta: {
        title: '用户管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: '/webhook',
      name: 'WebhookManagement',
      component: () => import('@/views/webhook/index.vue'),
      meta: {
        title: 'webhook管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: '/settings',
      name: 'SystemSettings',
      component: () => import('@/views/system/settings/index.vue'),
      meta: {
        title: '系统设置',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
    {
      path: '/api',
      name: 'ApiManagement',
      component: () => import('@/views/system/api/index.vue'),
      meta: {
        title: 'api管理',
        showLink: true,
        roles: ['admin', 'common'],
      },
    },
  ] satisfies RouteRecordRaw []
