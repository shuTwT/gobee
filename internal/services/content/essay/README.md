# Essay Service

## 功能说明

说说(即刻短文)模块的业务逻辑层，负责处理说说的数据操作和业务逻辑。

## 方法列表

### CreateEssay
创建一个新的说说

```go
func (s *EssayServiceImpl) CreateEssay(ctx context.Context, req *model.EssayCreateReq) (*ent.Essay, error)
```

**参数**:
- `ctx`: context.Context - 上下文
- `req`: *model.EssayCreateReq - 创建请求

**返回**:
- `*ent.Essay`: 创建的说说实体
- `error`: 错误信息

### UpdateEssay
更新指定的说说

```go
func (s *EssayServiceImpl) UpdateEssay(ctx context.Context, id int, req *model.EssayUpdateReq) error
```

**参数**:
- `ctx`: context.Context - 上下文
- `id`: int - 说说ID
- `req`: *model.EssayUpdateReq - 更新请求

**返回**:
- `error`: 错误信息

### GetEssay
获取指定的说说

```go
func (s *EssayServiceImpl) GetEssay(ctx context.Context, id int) (*ent.Essay, error)
```

**参数**:
- `ctx`: context.Context - 上下文
- `id`: int - 说说ID

**返回**:
- `*ent.Essay`: 说说明体
- `error`: 错误信息

### GetEssayPage
获取说说分页列表

```go
func (s *EssayServiceImpl) GetEssayPage(ctx context.Context, page, pageSize int) ([]*ent.Essay, int, error)
```

**参数**:
- `ctx`: context.Context - 上下文
- `page`: int - 页码
- `pageSize`: int - 每页数量

**返回**:
- `[]*ent.Essay`: 说说明体列表
- `int`: 总数
- `error`: 错误信息

### DeleteEssay
删除指定的说说

```go
func (s *EssayServiceImpl) DeleteEssay(ctx context.Context, id int) error
```

**参数**:
- `ctx`: context.Context - 上下文
- `id`: int - 说说ID

**返回**:
- `error`: 错误信息

## 数据模型

### Essay
说说明体包含以下字段:
- `id`: int - ID
- `content`: string - 内容
- `draft`: bool - 是否为草稿
- `images`: []string - 图片列表
- `created_at`: time.Time - 创建时间
- `updated_at`: time.Time - 更新时间
- `user_id`: int - 用户ID

## 使用示例

```go
// 创建说说
essay, err := essayService.CreateEssay(ctx, &model.EssayCreateReq{
    Content: "这是一条说说",
    Draft:   false,
    Images:  []string{"image1.jpg", "image2.jpg"},
})

// 获取分页列表
essays, total, err := essayService.GetEssayPage(ctx, 1, 10)

// 更新说说
err := essayService.UpdateEssay(ctx, 1, &model.EssayUpdateReq{
    ID:      1,
    Content: "更新后的内容",
    Draft:   false,
})

// 删除说说
err := essayService.DeleteEssay(ctx, 1)
```
