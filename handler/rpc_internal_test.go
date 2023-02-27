package handler

import (
	"context"
	"testing"

	hz "google.golang.org/grpc/health/grpc_health_v1"
)

func TestInternal(t *testing.T) {
	req := &hz.HealthCheckRequest{
		Service: m.baseCfg.Services.ServiceCode,
	}

	_, err := m.HealthCheck(context.TODO(), req)
	if err != nil {
		t.Errorf("HealthCheck test fail: %v", err)
	}
}
