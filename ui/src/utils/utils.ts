export function useStorageLocal(){

  function setItem<T>(key:string,value:T):void{
    localStorage.setItem(key,JSON.stringify(value))
  }
  function getItem<T>(key:string):T{
    return JSON.parse(localStorage.getItem(key) as string) as T
  }
  function removeItem(key:string):void{
    localStorage.removeItem(key)
  }

  function clear():void{
    localStorage.clear()
  }
  return {
    setItem,
    getItem,
    removeItem,
    clear
  }
}

/**
 * 判断字符串是否为合法链接（URL）
 * @param str 待检测的字符串
 * @returns 是链接则返回 true，否则 false
 */
export function isUrl(str: string): boolean {
  // 正则表达式：覆盖常见 URL 格式
  const urlRegex = /^(https?:\/\/|ftp:\/\/|mailto:)?(([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}|localhost|(\d{1,3}\.){3}\d{1,3})(:\d{1,5})?(\/[^\s]*)?$/;

  // 去除字符串前后空格后检测
  return urlRegex.test(str.trim());
}

