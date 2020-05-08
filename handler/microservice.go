// Code generated by grpc-kit-cli. DO NOT EDIT.

package handler

import (
	"github.com/grpc-kit/cfg"
	"github.com/grpc-kit/pkg/rpc"
	"github.com/sirupsen/logrus"

	"github.com/test/apilogin/modeler"
)

// Microservice 该微服务的结构
type Microservice struct {
	code    string                  // 服务代码
	server  *rpc.Server             // 服务调用
	logger  *logrus.Entry           // 全局日志
	baseCfg *cfg.LocalConfig        // 基础配置
	thisCfg *modeler.IndependentCfg // 个性配置
}

// NewMicroservice 全局只实例化一次
func NewMicroservice(lc *cfg.LocalConfig) (*Microservice, error) {
	// 根据本地配置完成初始化
	if err := lc.Init(); err != nil {
		return nil, err
	}

	m := &Microservice{
		code:    lc.Services.ServiceCode,
		logger:  lc.GetLogger(),
		baseCfg: lc,
		thisCfg: &modeler.IndependentCfg{},
	}

	if err := lc.GetIndependent(m.thisCfg); err != nil {
		return m, err
	}

	c := rpc.NewConfig(m.logger)
	c.GRPCAddress = lc.Services.GRPCAddress
	c.HTTPAddress = lc.Services.HTTPAddress

	m.server = rpc.NewServer(c, lc.GetUnaryInterceptor(m.privateUnaryServerInterceptor()...))

	// 其他私有的扩展操作
	if err := m.privateExtended(); err != nil {
		return m, err
	}

	return m, nil
}