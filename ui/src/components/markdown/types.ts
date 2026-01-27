import type { HeadList, SaveEvent } from "md-editor-v3"

export interface MdPreviewProps{
    /**
     * 输入框内容
     */
    value:string
}

export interface MdPreviewEmits{
    /**
     * 内容改变时触发
     */
    onChange:[value:string]
    /**
     * html变化回调时间，用于获取预览 html 代码
     */
    onHtmlChanged:[value:string]
    /**
     * 动态获取 markdown 目录
     */
    onGetCatalog:[list:HeadList[]]
    /**
     * 重新挂载时间
     */
    onRemount:[]
}

export interface MdEditorEmits extends MdPreviewEmits{
    /**
     * 保存事件,快捷键与按钮均会触发
     */
    onSave:[value:string,h:Promise<string>]
    onUploadImg:[files:Array<File>, callback: (urls: string[] | { url: string; alt: string; title: string }[])=>void]
    /**
     * 捕获执行错误事件,已将 InnerError 处理
     */
    onError:[error:Error]
    /**
     * 输入框失去焦点事件
     */
    onBlur:[event:FocusEvent]
    /**
     * 输入框获得焦点事件
     */
    onFocus:[event:FocusEvent]
    /**
     * 输入框输入事件
     */
    onInput:[event:Event]
    /**
     * 拖放内容事件
     */
    onDrop:[event:DragEvent]
    /**
     * 调整输入框宽度事件
     */
    onInputBoxWidthChange:[width:string]
}