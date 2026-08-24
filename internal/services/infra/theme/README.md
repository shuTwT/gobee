# 主题管理 Service（infra/theme）

主题系统的业务逻辑层：主题 zip 包的解析与落盘、ent `themes` 表读写、启用互斥控制，以及主题设置（setting.yaml 表单化）的读写。

## 依赖

- `*ent.Client`：数据访问
- `system/setting.SettingService`：主题配置值存储（setting 表，key 为 `theme:{主题name}`）

## 核心能力

### 主题生命周期

- `CreateTheme` 按 internal/static/external 分发；internal/static 从 zip 根目录读取并校验 `theme.yaml` 与 `setting.yaml`，解压到 `./data/themes/{name}/` 后入库
- `EnableTheme` 启用前先禁用所有其他主题（全局互斥）；已启用主题禁止删除

### 主题设置

- `GetThemeSetting(id)`：解析主题目录下的 `setting.yaml` 得到表单定义，叠加 setting 表已存值（缺失项回填 schema 默认值），返回给后台渲染设置表单
- `SaveThemeSetting(id, values)`：按当前 schema 过滤未知字段（主题升级后自动清理孤儿值）、校验必填项后，序列化存入 setting 表
- `GetThemeSettingValues(theme)`：前台渲染用，返回生效配置值，由 `frontend_router.go` 注入模板上下文 `Config`

### setting.yaml 格式

```yaml
type: setting
forms:
  - group: base        # 分组 key，一个分组对应后台设置弹窗的一个 Tab
    label: 基础设置
    formSchema:
      - type: text     # text/textarea/number/select/radio/switch/color/date/secret
        name: site_title
        label: 站点标题
        default: 星屑
        required: true # 可选属性：placeholder/help/min/max/options
```

校验规则（创建主题时执行，`validateSettingSchema`）：`type` 必须为空或 `setting`；分组 `group`/`label` 非空；字段 `name`/`label` 非空且 `name` 跨分组唯一；`type` 必须在允许集合内；`select`/`radio` 必须提供 `options`。

### 设计要点

- schema 不入库，每次从主题目录实时读取：主题升级重传后旧值保留，字段删除后保存时自动过滤
- 外部主题（external）无 `setting.yaml`，返回空表单与空值
- setting 表中值损坏时不阻断设置页打开（按空值处理，保存时覆盖）
