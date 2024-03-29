package handler

import (
	"context"
	"testing"

	pb "github.com/mingqing/apilogin/api/mingqing/apilogin/v1"
)

func TestDemo(t *testing.T) {
	req := &pb.DemoRequest{
		Uuid: "99feafb5-bed6-4daf-927a-69a2ab80c485",
	}

	resp, err := m.Demo(context.TODO(), req)
	if err != nil {
		t.Errorf("Demo test fail: %v", err)
	}

	if resp.Pong.Pong.Name != "grpc-kit" {
		t.Errorf("the expected name is: %v, but: %v", "grpc-kit", resp.Pong.Pong.Name)
	}
}
