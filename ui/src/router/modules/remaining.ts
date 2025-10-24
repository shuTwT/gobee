import BaseLayout from '@/layout/index.vue'

export default  [
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/login/index.vue'),
    },
    {
      path: '/',
      name: 'System',
      component: BaseLayout,
      children: [
        {
          path: '/',
          name: 'Home',
          component: () => import('@/views/home/index.vue'),
          meta:{
            title:'首页',
            showLink:true
          }
        },
        {
          path: '/album',
          name: 'AlbumManagement',
          component: () => import('@/views/album/index.vue'),
          meta:{
            title:'相册管理',
            showLink:true
          }
        },
        {
          path: '/comment',
          name: 'CommentManagement',
          component: () => import('@/views/comment/index.vue'),
          meta:{
            title:'评论管理',
            showLink:true
          }
        },
        {
          path: '/file',
          name: 'FileManagement',
          component: () => import('@/views/file/index.vue'),
          meta:{
            title:'文件管理',
            showLink:true
          }
        },
        {
          path: '/flink',
          name: 'FlinkManagement',
          component: () => import('@/views/flink/index.vue'),
          meta:{
            title:'flink管理',
            showLink:true
          }
        },
        {
          path: '/model-content',
          name: 'ModelContentManagement',
          component: () => import('@/views/model-content/index.vue'),
          meta:{
            title:'模型内容管理',
            showLink:true
          }
        },
        {
          path: '/model-schema',
          name: 'ModelSchemaManagement',
          component: () => import('@/views/model-schema/index.vue'),
          meta:{
            title:'模型内容管理',
            showLink:true
          }
        },
        {
          path: '/pay-channel',
          name: 'PayChannelManagement',
          component: () => import('@/views/pay-channel/index.vue'),
          meta:{
            title:'支付渠道管理',
            showLink:true
          }
        },
        {
          path: '/pay-order',
          name: 'PayOrderManagement',
          component: () => import('@/views/pay-order/index.vue'),
          meta:{
            title:'支付订单管理',
            showLink:true
          }
        },
        {
          path: '/user',
          name: 'UserManagement',
          component: () => import('@/views/system/user/index.vue'),
          meta:{
            title:'用户管理',
            showLink:true
          }
        },
        {
          path: '/webhook',
          name: 'WebhookManagement',
          component: () => import('@/views/webhook/index.vue'),
          meta:{
            title:'webhook管理',
            showLink:true
          }
        },
        {
          path: '/settings',
          name: 'SystemSettings',
          component: () => import('@/views/system/settings/index.vue'),
          meta:{
            title:'系统设置',
            showLink:true
          }
        },
        {
          path: '/api',
          name: 'ApiManagement',
          component: () => import('@/views/system/api/index.vue'),
          meta:{
            title:'api管理',
            showLink:true
          }
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
          component: () => import('@/views/error/404.vue'),
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
          component: () => import('@/views/error/403.vue'),
        },
      ],
    },
    {
    path: '/initialize',
    component: () => import('@/views/initialize/index.vue'),
  }
  ]
