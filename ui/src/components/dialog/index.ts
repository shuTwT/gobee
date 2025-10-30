import type { DialogOptions } from "./types"
import {useTimeoutFn} from "@vueuse/core"

const dialogStore = ref<Array<DialogOptions>>([])

/** 打开弹窗 */
const addDialog = (options:DialogOptions)=>{
  const open =()=>dialogStore.value.push(Object.assign({
    width:'600px'
  },options,{visible:true}))

  open()
}

const closeDialog = (options:DialogOptions,index:number,args?:any)=>{
  dialogStore.value[index].visible = false
  useTimeoutFn(()=>{
    dialogStore.value.splice(index,1)
  },200)
}

const updateDialg = (value:any,key='title',index=0)=>{
  dialogStore.value[index][key] = value
}

const closeAllDialog = ()=>{
  dialogStore.value = []
}

export {
  dialogStore,
  addDialog,
  closeDialog,
  updateDialg,
  closeAllDialog,
}
