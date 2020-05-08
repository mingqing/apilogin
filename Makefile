.PHONY: precheck clean version proto

include scripts/env

# 全局通用变量
GO              := go
GORUN           := ${GO} run
GOPATH          := $(shell ${GO} env GOPATH)
GOARCH          ?= $(shell ${GO} env GOARCH)
GOBUILD         := ${GO} build
WORKDIR         := $(shell pwd)
GOHOSTOS        := $(shell ${GO} env GOHOSTOS)
SHORTNAME       := $(shell basename $(shell pwd))
NAMESPACE       ?= $(shell basename $(shell cd ../;pwd))

# 自动化版本号
GIT_COMMIT      := $(shell git rev-parse HEAD 2>/dev/null)
GIT_BRANCH      := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
BUILD_DATE      := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
COMMIT_DATE     := $(shell git --no-pager log -1 --format='%ct' 2>/dev/null)
PREFIX_VERSION  := $(shell ./scripts/version.sh prefix)
RELEASE_VERSION ?= $(shell ./scripts/version.sh release)
BUILD_LD_FLAGS  := "-X 'github.com/grpc-kit/pkg/version.CliVersion=${CLI_VERSION}' \
                -X 'github.com/grpc-kit/pkg/version.GitCommit=${GIT_COMMIT}' \
                -X 'github.com/grpc-kit/pkg/version.GitBranch=${GIT_BRANCH}' \
                -X 'github.com/grpc-kit/pkg/version.BuildDate=${BUILD_DATE}' \
                -X 'github.com/grpc-kit/pkg/version.CommitUnixTime=${COMMIT_DATE}' \
                -X 'github.com/grpc-kit/pkg/version.ReleaseVersion=${RELEASE_VERSION}'"

# 构建Docker容器变量
IMAGE_FROM      ?= scratch
IMAGE_HOST      ?= hub.docker.com
BUILD_GOOS      ?= $(shell ${GO} env GOOS)

precheck:
	@echo ">> precheck environment"
	@./scripts/precheck.sh

run: generate
	@${GORUN} -ldflags ${BUILD_LD_FLAGS} cmd/server/main.go -c config/app-dev-local.toml

build: clean generate
	@mkdir build
	@mkdir build/deploy
	@GOOS=${BUILD_GOOS} ${GOBUILD} -ldflags ${BUILD_LD_FLAGS} -o build/service cmd/server/main.go

clean:
	@echo ">> clean build"
	@rm -rf build/

generate: precheck
	@echo ">> generation release version"
	@./scripts/version.sh update

	@echo ">> generation assets to static code"
	@${GO} generate api/generate.go >> /dev/null

	@echo ">> generation code from proto files"
	@./scripts/genproto.sh

build-docker: build
	@echo ">> build docker"

	@IMAGE_FROM=${IMAGE_FROM} \
	IMAGE_HOST=${IMAGE_HOST} \
	NAMESPACE=${NAMESPACE} \
	SHORTNAME=${SHORTNAME} \
	RELEASE_VERSION=${RELEASE_VERSION} \
	./scripts/docker.sh