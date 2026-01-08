import type { RouteRecordRaw } from "vue-router";

export default [
    {
        path: "/infra",
        redirect: "/infra/file",
        meta: {
            title: "基础设施",
            showLink: true,
            roles: ['admin', 'common'],
            rank:99
        },
        children: [
            {
                path: '/file',
                name: 'FileManagement',
                component: () => import('@/views/infra/file/index.vue'),
                meta: {
                    title: '文件管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/storage-strategy',
                name: 'StorageStrategyManagement',
                component: () => import('@/views/infra/storage-strategy/index.vue'),
                meta: {
                    title: '存储策略管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            }
        ]
    }

] satisfies RouteRecordRaw[]
