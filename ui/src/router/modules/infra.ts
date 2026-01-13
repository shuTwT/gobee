import type { RouteRecordRaw } from "vue-router";

export default [
    {
        path: "/infra",
        name:"InfrastructureManagement",
        redirect: "/infra/file",
        meta: {
            title: "基础设施",
            showLink: true,
            roles: ['admin', 'common'],
            rank:99
        },
        children: [
            {
                path: '/infra/file',
                name: 'FileManagement',
                component: () => import('@/views/infra/file/index.vue'),
                meta: {
                    title: '文件管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/infra/storage-strategy',
                name: 'StorageStrategyManagement',
                component: () => import('@/views/infra/storage-strategy/index.vue'),
                meta: {
                    title: '存储策略管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/infra/schedule-job',
                name: 'ScheduleJobManagement',
                component: () => import('@/views/infra/schedulejob/index.vue'),
                meta: {
                    title: '定时任务管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/infra/migration',
                name: 'MigrationManagement',
                component: () => import('@/views/infra/migration/index.vue'),
                meta: {
                    title: '迁移',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            },
            {
                path: '/infra/visit-log',
                name: 'VisitLogManagement',
                component: () => import('@/views/infra/visit-log/index.vue'),
                meta: {
                    title: '访问日志',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            }
        ]
    }

] satisfies RouteRecordRaw[]
