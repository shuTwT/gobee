import Cookies from "js-cookie"

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
  let expires = 0;
  const {accessToken,refreshToken} = data
  const {}
}
