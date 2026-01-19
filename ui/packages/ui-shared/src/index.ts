import type{RouteRecordRaw} from 'vue-router'
import type{Component} from 'vue'

export interface PluginModule{
    components?:Record<string,Component>
    routes?:RouteRecordRaw[]
}

export interface ExtensionPoint{
    [key:string]:any
}
export function definePlugin(plugin:PluginModule):PluginModule{
    return plugin
}