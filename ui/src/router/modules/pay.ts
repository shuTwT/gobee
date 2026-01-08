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
            rank: 40
        },
        children: [
            {
                path: '/pay-order',
                name: 'PayOrderManagement',
                component: () => import('@/views/pay/pay-order/index.vue'),
                meta: {
                    title: '支付订单管理',
                    showLink: true,
                    roles: ['admin', 'common'],
                },
            }
        ]
    }
] satisfies RouteRecordRaw[]