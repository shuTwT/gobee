import { usePermissionStore, usePermissionStoreHook } from "@/stores/modules/permission";
import type { RouteComponent } from "vue-router";
import { cloneDeep } from "lodash-es"
import router from ".";
import { getAsyncRoutes } from "@/api/routes";

function isAllEmpty(value:any){
  return value === undefined || value === null || value === ''
}

function handRank(routeInfo: any) {
  const { name, path, parentId, meta } = routeInfo;
  return isAllEmpty(parentId)
    ? isAllEmpty(meta?.rank) ||
      (meta?.rank === 0 && name !== "Home" && path !== "/")
      ? true
      : false
    : false;
}

/** 按照路由中meta下的rank等级升序来排序路由 */
function ascending(arr: any[]) {
  arr.forEach((v, index) => {
    // 当rank不存在时，根据顺序自动创建，首页路由永远在第一位
    if (handRank(v)) v.meta.rank = index + 2;
  });
  return arr.sort(
    (a: { meta: { rank: number } }, b: { meta: { rank: number } }) => {
      return a?.meta.rank - b?.meta.rank;
    }
  );
}

/** 过滤meta中showLink为false的菜单 */
function filterTree(data: RouteComponent[]) {
  const newTree = (cloneDeep(data) as any).filter(
    (v: { meta: { showLink: boolean } }) => v.meta?.showLink !== false
  );
  newTree.forEach(
    (v: { children:any }) => v.children && (v.children = filterTree(v.children))
  );
  return newTree;
}


function handleAsyncRoutes(routeList: unknown[]){
  if(routeList.length === 0){
    usePermissionStoreHook().handleWholeMenus(routeList)
  }else{
    usePermissionStoreHook().handleWholeMenus(routeList)
  }
  addPathMatch()
}

function addPathMatch() {
  if (!router.hasRoute("pathMatch")) {
    router.addRoute({
      path: "/:pathMatch(.*)",
      name: "pathMatch",
      redirect: "/error/404"
    });
  }
}

function initRouter(){
  return new Promise(resolve=>{
    getAsyncRoutes().then(({data})=>{
      handleAsyncRoutes(data)
      resolve(router)
    })
  })
}

export {
  ascending,
  filterTree
}
