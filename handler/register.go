// Code generated by "grpc-kit-cli/0.2.4". DO NOT EDIT.

package handler

import (
	"context"
	"net/http"

	pb "github.com/mingqing/apilogin/api/mingqing/apilogin/v1"
	"github.com/mingqing/apilogin/public/doc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Register 用于服务启动前环境准备
func (m *Microservice) Register(ctx context.Context) error {
	pb.RegisterMingqingApiloginServer(m.server.Server(), m)
	grpc_health_v1.RegisterHealthServer(m.server.Server(), health.NewServer())

	// 注册服务信息
	mux, err := m.baseCfg.Register(ctx, pb.RegisterMingqingApiloginHandlerFromEndpoint)
	if err != nil {
		return err
	}

	// 注册API文档
	mux.Handle("/openapi-spec/", http.FileServer(http.FS(doc.Assets)))

	// 这里添加其他自定义实现
	m.privateHTTPHandle(mux)

	// 注册HTTP网关
	if err := m.server.RegisterGateway(mux); err != nil {
		return err
	}

	// 开启gRPC与HTTP服务
	if err := m.server.StartBackground(); err != nil {
		return err
	}

	return nil
}
