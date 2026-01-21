# License Service

授权管理模块的服务层，负责处理授权相关的业务逻辑。

## 功能

- 创建授权
- 查询授权分页列表
- 查询单个授权详情
- 更新授权信息
- 删除授权
- 验证授权

## 方法

### CreateLicense
创建新的授权记录。

参数:
- ctx: 上下文
- domain: 域名
- licenseKey: 授权密钥
- customerName: 客户名称
- expireDate: 过期日期

返回: 创建的授权实体

### ListLicensePage
查询授权分页列表。

参数:
- ctx: 上下文
- page: 页码
- size: 每页数量

返回: 总数和授权列表

### QueryLicense
查询单个授权详情。

参数:
- ctx: 上下文
- id: 授权ID

返回: 授权实体

### UpdateLicense
更新授权信息。

参数:
- ctx: 上下文
- id: 授权ID
- domain: 域名
- licenseKey: 授权密钥
- customerName: 客户名称
- expireDate: 过期日期
- status: 状态 (1-有效, 2-过期, 3-禁用)

返回: 更新后的授权实体

### DeleteLicense
删除授权记录。

参数:
- ctx: 上下文
- id: 授权ID

返回: 错误信息

### VerifyLicense
验证授权是否有效。

参数:
- ctx: 上下文
- domain: 域名

返回: 授权实体（如果有效）或 nil（如果无效）

## 业务逻辑

### 授权验证逻辑
1. 根据域名查询授权记录
2. 检查授权是否存在
3. 检查授权是否过期
4. 如果过期，自动更新状态为"过期"
5. 检查授权状态是否为"有效"
6. 返回验证结果

### 状态定义
- 1: 有效
- 2: 过期
- 3: 禁用

## 错误处理

所有方法都包含完善的错误处理：

- 数据库操作失败：返回错误
- 授权不存在：返回错误
- 授权过期：返回 nil
- 授权被禁用：返回 nil
