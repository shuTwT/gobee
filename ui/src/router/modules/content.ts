import type { RouteRecordRaw } from 'vue-router'

export default [
    {
        path: "/content",
        name:"ContentManagement",
        redirect: "/content/post",
        meta: {
            title: '内容管理',
            showLink: true,
            roles: ['admin', 'common'],
            rank: 10
        },
        children: [
            {
                path: '/content/post',
                name: 'PostManagement',
                component: () => import('@/views/content/post/index.vue'),
                meta: {
                    title: '文章管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/content/post/edit',
                name: 'PostEditor',
                component: () => import('@/views/content/post/editor/index.vue'),
                meta: {
                    title: '文章编辑',
                    showLink: false,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/content/album',
                name: 'AlbumManagement',
                component: () => import('@/views/content/album/index.vue'),
                meta: {
                    title: '相册管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/content/comment',
                name: 'CommentManagement',
                component: () => import('@/views/comment/index.vue'),
                meta: {
                    title: '评论管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/content/flink',
                name: 'FlinkManagement',
                component: () => import('@/views/content/flink/index.vue'),
                meta: {
                    title: '友链管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/content/friend-circle',
                name: 'FriendCircleRecord',
                component: () => import('@/views/friend-circle/index.vue'),
                meta: {
                    title: '朋友圈',
                    showLink: true,
                    roles: ['admin'],
                },
            },
        ]
    }
] satisfies RouteRecordRaw[]