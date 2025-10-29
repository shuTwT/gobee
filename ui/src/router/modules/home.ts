import type { RouteRecordRaw } from "vue-router";
import Layout from '@/layout/index.vue'

export default [
  {
    path:"/",
    name:"Home",
    component:Layout,
    redirect:"/welcome",
    meta:{
      title:"首页",
      rank:0,
      roles:["admin","common"]
    },
    children:[
      {
        path:"/welcome",
        name:"Welcome",
        component:()=>import("@/views/home/index.vue"),
        meta:{
          title:"首页",
          showLink:true,
          roles:["admin","common"]
        }
      }
    ]
  }
] satisfies RouteRecordRaw []
