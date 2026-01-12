import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: '/mall',
    name: 'MallManagement',
    redirect: '/mall/product',
    meta: {
      title: '商城管理',
      showLink: true,
      roles: ['admin'],
      rank: 20,
    },
    children: [
      {
        path: '/mall/product',
        name: 'ProductList',
        component: () => import('@/views/mall/product/index.vue'),
        meta: {
          title: '商品列表',
          showLink: true,
          roles: ['admin'],
          rank: 21,
        },
      },
      {
        path: '/mall/member',
        name: 'MemberList',
        component: () => import('@/views/mall/member/index.vue'),
        meta: {
          title: '会员列表',
          showLink: true,
          roles: ['admin'],
          rank: 22,
        },
      },
      {
        path: '/member/level',
        name: 'MemberLevelManagement',
        component: () => import('@/views/mall/member-level/index.vue'),
        meta: {
          title: '会员等级管理',
          showLink: true,
          roles: ['admin'],
          rank: 23,
        },
      },
      {
        path: '/mall/coupon',
        name: 'CouponManagement',
        component: () => import('@/views/mall/coupon/index.vue'),
        meta: {
          title: '优惠券管理',
          showLink: true,
          roles: ['admin'],
          rank: 24,
        },
      },
      {
        path: '/mall/coupon-usage',
        name: 'CouponUsageManagement',
        component: () => import('@/views/mall/couponUsage/index.vue'),
        meta: {
          title: '优惠券使用记录',
          showLink: true,
          roles: ['admin'],
          rank: 25,
        },
      },
      {
        path: '/mall/pay-order',
        name: 'PayOrderManagement',
        component: () => import('@/views/mall/pay-order/index.vue'),
        meta: {
          title: '支付订单管理',
          showLink: true,
          roles: ['admin', 'common'],
          rank: 26,
        },
      },
    ],
  },
] satisfies RouteRecordRaw[]
