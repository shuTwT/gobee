# License Client

授权管理客户端模块，提供机器码生成、授权安装、授权验证和清除授权功能。

## 功能

- 生成机器码：基于系统信息生成唯一的机器标识
- 安装授权：将授权信息保存到本地
- 验证授权：检查本地授权是否有效
- 清除授权：删除本地授权信息

## 使用方法

### 生成机器码

```go
import "github.com/shuTwT/gobee/internal/infra/license"

machineCode, err := license.GenerateMachineCode()
if err != nil {
    log.Fatal(err)
}
fmt.Println("机器码:", machineCode)
```

### 安装授权

```go
import "github.com/shuTwT/gobee/internal/infra/license"

err := license.InstallLicense(
    "machine-code",
    "license-key",
    "customer-name",
    "2025-12-31T23:59:59Z",
)
if err != nil {
    log.Fatal(err)
}
fmt.Println("授权安装成功")
```

### 验证授权

```go
import "github.com/shuTwT/gobee/internal/infra/license"

licenseInfo, err := license.VerifyLicense()
if err != nil {
    log.Fatal("授权验证失败:", err)
}
fmt.Printf("授权有效: 客户=%s, 过期时间=%s\n", licenseInfo.CustomerName, licenseInfo.ExpireDate)
```

### 清除授权

```go
import "github.com/shuTwT/gobee/internal/infra/license"

err := license.ClearLicense()
if err != nil {
    log.Fatal(err)
}
fmt.Println("授权已清除")
```

## 授权文件位置

授权信息保存在用户主目录下的 `.gobee/license.json` 文件中。

- Linux/macOS: `~/.gobee/license.json`
- Windows: `C:\Users\<username>\.gobee\license.json`

## 授权文件格式

```json
{
  "machine_code": "机器码",
  "license_key": "授权密钥",
  "customer_name": "客户名称",
  "expire_date": "2025-12-31T23:59:59Z"
}
```

## 机器码生成逻辑

机器码基于以下信息生成：
- 主机名
- 操作系统类型
- 系统架构
- CPU 核心数
- Go 版本

使用 MD5 哈希算法生成唯一的机器码。

## 授权验证逻辑

1. 读取本地授权文件
2. 解析授权信息
3. 生成当前机器的机器码
4. 比对机器码是否匹配
5. 检查授权是否过期
6. 返回验证结果

## 错误处理

所有函数都包含完善的错误处理：

- 机器码生成失败：返回错误
- 授权文件不存在：返回错误
- 授权文件格式错误：返回错误
- 机器码不匹配：返回错误
- 授权已过期：返回错误
- 清除授权失败：返回错误

## 使用场景

### 应用启动时验证授权

```go
func main() {
    licenseInfo, err := license.VerifyLicense()
    if err != nil {
        log.Fatal("授权验证失败:", err)
    }
    log.Printf("授权有效: 客户=%s", licenseInfo.CustomerName)
    
    // 启动应用
    startApplication()
}
```

### 获取机器码用于授权申请

```go
func getMachineCode() string {
    machineCode, err := license.GenerateMachineCode()
    if err != nil {
        log.Fatal(err)
    }
    return machineCode
}
```

### 重新授权

```go
func reinstallLicense(newLicenseKey string) {
    // 清除旧授权
    license.ClearLicense()
    
    // 安装新授权
    machineCode, _ := license.GenerateMachineCode()
    err := license.InstallLicense(
        machineCode,
        newLicenseKey,
        "customer-name",
        "2026-12-31T23:59:59Z",
    )
    if err != nil {
        log.Fatal(err)
    }
}
```
