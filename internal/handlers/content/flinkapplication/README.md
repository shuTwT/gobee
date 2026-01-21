# 友链申请管理模块

## 功能说明

友链申请管理模块用于管理用户提交的友链申请，包括新增和修改申请的审批流程。

## 主要功能

- 创建友链申请：前台用户可以提交友链申请，包括新增和修改两种类型
- 查询友链申请列表：管理员可以查看所有友链申请，支持分页和状态筛选
- 查询单个友链申请：管理员可以查看申请的详细信息
- 审批友链申请：管理员可以批准或拒绝友链申请

## 申请类型

- `create`：新增友链申请
- `update`：修改现有友链申请

## 审批状态

- `0`：待审批
- `1`：已通过
- `2`：已拒绝

## 字段说明

| 字段名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| website_url | string | 是 | 网站链接 |
| application_type | string | 是 | 申请类型（create/update） |
| website_name | string | 是 | 网站名称 |
| website_logo | string | 是 | 网站 logo |
| website_description | string | 是 | 网站描述 |
| contact_email | string | 是 | 联系邮箱 |
| snapshot_url | string | 否 | 网页快照 |
| original_website_url | string | 否 | 原网站链接（修改类型时必填） |
| modification_reason | string | 否 | 修改原因（修改类型时必填） |
| status | int | 否 | 审批状态（0-待审批, 1-已通过, 2-已拒绝） |
| reject_reason | string | 否 | 拒绝原因 |
