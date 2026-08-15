// Package capability 定义插件的通用能力调用协议（宿主与插件共享）。
// 插件实现 Capability 接口暴露业务能力，宿主通过统一的 Call 调用，
// 新增插件无需修改宿主代码。
package capability

import (
	"context"
	"errors"

	"github.com/hashicorp/go-plugin"
	pb "github.com/shuTwT/hoshikuzu/pkg/plugin/capability/proto"
	"google.golang.org/grpc"
)

// Capability 插件通用能力接口：method 为插件自定义方法名，
// params/result 均为 JSON 原文（字节），由插件侧负责解析与序列化
type Capability interface {
	Call(ctx context.Context, method string, params []byte) ([]byte, error)
}

// CapabilityPlugin 能力协议载体，实现 hashicorp/go-plugin 框架的 Plugin 接口
type CapabilityPlugin struct {
	plugin.Plugin
	Impl Capability
}

// GRPCServer 插件侧运行：将能力实现注册到 gRPC 服务端
func (p *CapabilityPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterCapabilityServiceServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

// GRPCClient 宿主侧运行：创建 gRPC 客户端并封装为 Capability 接口
func (p *CapabilityPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: pb.NewCapabilityServiceClient(c)}, nil
}

// GRPCClient 宿主侧客户端，实现 Capability 接口
type GRPCClient struct {
	client pb.CapabilityServiceClient
}

func (c *GRPCClient) Call(ctx context.Context, method string, params []byte) ([]byte, error) {
	resp, err := c.client.Call(ctx, &pb.CallRequest{
		Method: method,
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}
	return resp.Result, nil
}

// GRPCServer 插件侧服务端，将 gRPC 请求转发给 Capability 实现
type GRPCServer struct {
	pb.UnimplementedCapabilityServiceServer
	Impl Capability
}

func (s *GRPCServer) Call(ctx context.Context, req *pb.CallRequest) (*pb.CallResponse, error) {
	result, err := s.Impl.Call(ctx, req.Method, req.Params)
	if err != nil {
		return &pb.CallResponse{Error: err.Error()}, nil
	}
	return &pb.CallResponse{Result: result}, nil
}

// 确保 GRPCClient 完全实现 Capability 接口（编译期校验）
var _ Capability = &GRPCClient{}
