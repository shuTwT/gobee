import type { ExtendProForm } from "pro-naive-ui";

export type CreateProSearchFormReturn<Values = any> = ExtendProForm<Values, {
    /**
     * 是否收起
     */
    collapsed: Ref<boolean>;
    /**
     * 切换收起
     * @param collapsed 传递了此参数，根据参数切换
     */
    toggleCollapse: (collapsed?: boolean) => void;
}>;
