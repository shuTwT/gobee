import type { ButtonProps, ModalProps, PopconfirmProps } from "naive-ui";

type EventType = "open"|"close"|"fullscreenCallBack"

type BtnClickDialog = {
  options?: DialogOptions;
  index?: number;
};
type BtnClickButton = {
  btn?: ButtonProps;
  index?: number;
};

interface ModalButtonProps extends ButtonProps {
  label:string
    /** 确定按钮的 `Popconfirm` 气泡确认框相关配置 */
  popconfirm?: PopconfirmProps;
  btnClick?:({
    dialog,button
  }:{
    dialog:BtnClickDialog,
    button:BtnClickButton
  })=>void
}

interface DialogOptions<P=Record<string,any>> extends ModalProps{
  props?:P
  visible?:boolean
  fullscreen?:boolean
  scroll?:boolean
  scrollbarHeight?:number|string
  /** 自定义底部按钮操作 */
  footerButtons?: Array<ModalButtonProps>;
  /** 点击确定按钮后是否开启 `loading` 加载动画 */
  sureBtnLoading?: boolean;
  /** 自定义内容渲染器 */
  contentRenderer: ({
    options,
    index
  }: {
    options: DialogOptions;
    index: number;
  }) => VNode | Component;
  /** 点击底部取消按钮的回调，会暂停 `Dialog` 的关闭. 回调函数内执行 `done` 参数方法的时候才是真正关闭对话框的时候 */
  beforeCancel?: (
    done: ()=>void,
    {
      options,
      index
    }: {
      options: DialogOptions;
      index: number;
    }
  ) => void;
  /** 点击底部确定按钮的回调，会暂停 `Dialog` 的关闭. 回调函数内执行 `done` 参数方法的时候才是真正关闭对话框的时候 */
  beforeSure?: (
    done: ()=>void,
    {
      options,
      index,
      closeLoading
    }: {
      options: DialogOptions;
      index: number;
      /** 关闭确定按钮的 `loading` 加载动画 */
      closeLoading: ()=>void;
    }
  ) => void;
  width?:number|string
  /** 在打开前获取数据 */
  fetchOpen?:()=>void
  [key:string]:any
}

export type {
  EventType,
  ModalButtonProps,
  DialogOptions
}
