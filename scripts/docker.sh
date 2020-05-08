#!/bin/sh

# 容器镜像版本号，去掉v开头
IMAGE_VERSION=${RELEASE_VERSION}
if test ${RELEASE_VERSION:0:1} = "v"; then
    IMAGE_VERSION=${RELEASE_VERSION:1}
fi

# 如未设置父镜像，默认为scratch
if test -z ${IMAGE_FROM}; then
    IMAGE_FROM=scratch
fi

# 生成的容器镜像地址
IMAGE_ADDR=${IMAGE_HOST}/${NAMESPACE}/${SHORTNAME}:${IMAGE_VERSION}

cp scripts/templates/Dockerfile ./

GOHOSTOS=$(go env GOHOSTOS)

if test ${GOHOSTOS} = "darwin"; then
    sed -i "" "s#{{IMAGE_FROM}}#${IMAGE_FROM}#g" Dockerfile
else
    sed -i "s#{{IMAGE_FROM}}#${IMAGE_FROM}#g" Dockerfile
fi

docker build -t ${IMAGE_ADDR} ./
#docker push ${IMAGE_ADDR}
echo "Now you can upload image: "docker push ${IMAGE_ADDR}""