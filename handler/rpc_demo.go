package handler

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogo/protobuf/types"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	kitApiV1 "github.com/grpc-kit/api/proto/v1"
	"github.com/grpc-kit/pkg/errors"
	pb "github.com/test/apilogin/api/proto/v1alpha1"
)

// Demo test
func (m Microservice) Demo(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	m.logger.Warnf("test demo warn: %v", "func Demo")

	result := &pb.DemoResponse{
		// GET /demo
		Content: []*kitApiV1.ExampleResponse{
			{Name: "grpc-kit-cli"},
			{Name: "grpc-kit-cfg"},
			{Name: "grpc-kit-pkg"},
			{Name: "grpc-kit-api"},
			{Name: "grpc-kit-web"},
			{Name: "grpc-kit-doc"},
		},
		Ping: &kitApiV1.ExampleResponse{},
		// POST /demo
		// GET /demo/{uuid}
		Pong: &pb.DemoResponse_Pong{
			Uuid: "99feafb5-bed6-4daf-927a-69a2ab80c485",
			Pong: &kitApiV1.ExampleResponse{},
		},
		// DELETE /demo/{uuid}
		Empty: &types.Empty{},
	}

	if req.Ping != nil {
		result.Ping.Name = req.Ping.Name
		result.Pong.Pong.Name = req.Ping.Name
	}

	if req.Uuid == "99feafb5-bed6-4daf-927a-69a2ab80c485" {
		result.Pong.Pong.Name = "grpc-kit"
	}

	return result, nil
}

// BaseAuth 验证base auth是否有效
func (m Microservice) BaseAuth(ctx context.Context, req *pb.BaseAuthRequest) (*pb.BaseAuthResponse, error) {
	authToken, err := grpcauth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(authToken)
	if err != nil {
		return nil, err
	}

	tmps := strings.Split(string(decoded), ":")
	if len(tmps) != 2 {
		return nil, errors.InvalidArgument(ctx).Err()
	}

	if tmps[0] != "hello" && tmps[1] != "world" {
		return nil, errors.Unauthenticated(ctx).Err()
	}

	result := &pb.BaseAuthResponse{
		LoginTime: time.Now().Format(time.RFC3339),
	}

	return result, nil
}

// JWTAuth 验证jwt是否有效
func (m Microservice) JWTAuth(ctx context.Context, req *pb.JWTAuthRequest) (*pb.JWTAuthResponse, error) {
	authToken, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	jwtToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if jwtToken == nil {
		return nil, errors.Unauthenticated(ctx).Err()
	}

	result := &pb.JWTAuthResponse{
		LoginTime: time.Now().Format(time.RFC3339),
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
		if err := claims.Valid(); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.Unauthenticated(ctx).Err()
	}

	return result, nil
}
