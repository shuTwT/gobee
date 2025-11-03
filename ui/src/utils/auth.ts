import { useUserStoreHook } from "@/stores/modules/user";
import Cookies from "js-cookie"
import { useStorageLocal } from "./utils";
import { every, includes, isEmpty, isString } from "lodash-es";

export interface DataInfo<T> {
  /** token */
  accessToken: string;
  /** `accessToken`的过期时间（时间戳） */
  expires: T;
  /** 用于调用刷新accessToken的接口时所需的token */
  refreshToken: string;
  /** 头像 */
  avatar?: string;
  /** 用户名 */
  username?: string;
  /** 昵称 */
  nickname?: string;
  /** 当前登录用户的角色 */
  roles?: Array<string>;
  /** 当前登录用户的按钮级别权限 */
  permissions?: Array<string>;
}

export const userKey = 'user-info'
export const tokenKey = 'authorized-token'
export const multipleTabsKey = 'multiple-tabs'

export function getToken(): DataInfo<number> | null{
  return Cookies.get(tokenKey)
  ?JSON.parse(Cookies.get(tokenKey)+'')
  :localStorage.getItem(tokenKey)?JSON.parse(localStorage.getItem(tokenKey)+'')
  :null
}

export function setToken(data:DataInfo<Date>){
  console.log(data)
  let expires = 0;
  const {accessToken,refreshToken} = data
  const {isRemembered,loginDay} = useUserStoreHook()
  expires = new Date(Number(data.expires)).getTime()
  console.log(expires)
  const cookieString = JSON.stringify({accessToken,refreshToken,expires})

  if(expires > 0){
    console.log("expires",(expires - Date.now()) / 86400000)
    Cookies.set(tokenKey, cookieString, {
      expires: (expires - Date.now()) / 86400000
    });
  }else{
    Cookies.set(tokenKey, cookieString);
  }

  Cookies.set(
    multipleTabsKey,
    "true",
    isRemembered
      ? {
          expires: loginDay
        }
      : {}
  );

  function setUserKey({avatar,username,nickname,roles,permissions}:Partial<DataInfo<number>>){
    useUserStoreHook().SET_AVATAR(avatar ??"")
    useUserStoreHook().SET_USERNAME(username??"")
    useUserStoreHook().SET_NICKNAME(nickname??"")
    useUserStoreHook().SET_ROLES(roles??[])
    useUserStoreHook().SET_PERMS(permissions??[])
    useStorageLocal().setItem(userKey,{avatar,username,nickname,roles,permissions})
  }

  if(data.username&&data.roles){
    const {username,roles} = data
    setUserKey({
      avatar:data?.avatar ?? '',
      username,
      nickname:data?.nickname??'',
      roles,
      permissions:data.permissions??[]
    })
  }else{
    const avatar = useStorageLocal().getItem<DataInfo<number>>(userKey)?.avatar ?? ''
    const username = useStorageLocal().getItem<DataInfo<number>>(userKey)?.username ?? ''
    const nickname = useStorageLocal().getItem<DataInfo<number>>(userKey)?.nickname ?? ''
    const roles = useStorageLocal().getItem<DataInfo<number>>(userKey)?.roles ?? []
    const permissions = useStorageLocal().getItem<DataInfo<number>>(userKey)?.permissions ?? []
    setUserKey({avatar,username,nickname,roles,permissions})
  }
}


export function removeToken(){
  Cookies.remove(tokenKey)
  Cookies.remove(multipleTabsKey)
  localStorage.removeItem(userKey)
}

/** 格式化token（jwt格式） */
export const formatToken = (token: string): string => {
  return "Bearer " + token;
};

/** 是否有按钮级别的权限（根据登录接口返回的`permissions`字段进行判断）*/
export const hasPerms = (value: string | Array<string>): boolean => {
  if (!value) return false;
  const allPerms = "*:*:*";
  const { permissions } = useUserStoreHook();
  if (!permissions) return false;
  if (permissions.length === 1 && permissions[0] === allPerms) return true;
  const isAuths = isString(value)
    ? permissions.includes(value)
    : isIncludeAllChildren(value, permissions);
  return isAuths ? true : false;
};

function isIncludeAllChildren(parent: Array<string>, child: Array<string>): boolean {
  // 子体为空数组时，默认“包含所有元素”（空数组无元素需要匹配）
  if (isEmpty(child)) return true;
  // 检查子体的每个元素是否都在母体中
  return every(child, (element) => includes(parent, element));
}
