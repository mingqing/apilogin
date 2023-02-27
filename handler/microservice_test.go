package handler

import (
	"testing"

	"github.com/grpc-kit/pkg/cfg"

	"github.com/mingqing/apilogin/modeler"
)

var m *Microservice

func init() {
	lc := &cfg.LocalConfig{
		Services: &cfg.ServicesConfig{
			ServiceCode: "apilogin.v1.mingqing",
			HTTPAddress: "127.0.0.1:8080",
			GRPCAddress: "127.0.0.1:10081",
		},
		Independent: modeler.IndependentCfg{},
	}

	s, err := NewMicroservice(lc)
	if err != nil {
		panic(err)
	}

	m = s
}

func TestMicroservice(t *testing.T) {
	t.Run("Main", func(t *testing.T) {
		t.Run("NotNil", testMicroserviceMainNotNil)
	})
}

func testMicroserviceMainNotNil(t *testing.T) {
	if m == nil {
		t.Error("microservice is nil")
	}
}
