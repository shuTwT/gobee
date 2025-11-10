package model

type ApiRoute struct {
	// api名称
	Name string `json:"name"`
	// api路径
	Path string `json:"path"`
	// 请求方法
	Method string `json:"method"`
	// 描述
	Desc string `json:"desc"`
	// 权限
	PermissionType string   `json:"permission_type"`
	Roles          []string `json:"roles"`
	// 分类
	Category []string `json:"category"`
	// 状态
	Status string `json:"status"`
}

var PermissionType = struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}{
	Public:  "public",
	Private: "private",
}
