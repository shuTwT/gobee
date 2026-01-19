package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Greeter interface {
	// 初始化
	Init() error
	// 启动
	Start() error
	// 停止
	Stop() error
}

type GreeterRPC struct {
	client *rpc.Client
}

type GreeterRPCServer struct {
	Impl Greeter
}

type GreeterPlugin struct {
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}
