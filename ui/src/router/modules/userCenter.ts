import type { RouteRecordRaw } from 'vue-router'

export default [
  {
    path: "/user-center",
    name:"UserCenter",
    component: () => import('@/views/user-center/index.vue'),
    meta: {
      showLink:false
    },
  }
] satisfies RouteRecordRaw []
