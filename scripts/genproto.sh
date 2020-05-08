#!/bin/sh

source scripts/env

# 生成*.pb.go文件
protoc -I./api/proto/${API_VERSION}/ \
        -I${GOPATH}/src \
        -I${GOPATH}/src/github.com/gogo/googleapis/ \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --gogo_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types:\
./api/proto/${API_VERSION}/ \
        api/proto/${API_VERSION}/*.proto

# 生成*.pb.gw.go与swagger接口文档
protoc -I./api/proto/${API_VERSION}/ \
        -I${GOPATH}/src \
        -I${GOPATH}/src/github.com/gogo/googleapis/ \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --grpc-gateway_out=allow_patch_feature=false,allow_repeated_fields_in_body=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types:\
./api/proto/${API_VERSION}/ \
        --swagger_out=logtostderr=true,allow_repeated_fields_in_body=true:./api/doc/openapi-spec/ \
        api/proto/${API_VERSION}/microservice.proto