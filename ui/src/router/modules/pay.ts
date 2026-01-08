import type { RouteRecordRaw } from 'vue-router'

export default [
    {
        path: '/pay',
        name: 'PayManagement',
        redirect:"/pay-order",
        meta: {
            title: '支付订单管理',
            showLink: true,
            roles: ['admin', 'common'],
        },
        children: [
            {
                path: '/pay-order',
                name: 'PayOrderManagement',
                component: () => import('@/views/pay-order/index.vue'),
                meta: {
                    title: '支付订单管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            }
        ]
    }
] satisfies RouteRecordRaw[]