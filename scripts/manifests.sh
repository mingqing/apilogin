#!/bin/bash

source scripts/env

if test -z $1; then
  echo "Usage:"
  echo "\t ./scripts/manifests.sh dev"
fi

GOHOSTOS=$(go env GOHOSTOS)
DEPLOY_ENV=$1

# 解决 linux 与 darwin 的 sed 存在的差异
if test ${GOHOSTOS} = "darwin"; then
  SED="sed -i ''"
else
  SED="sed -i"
fi

# 生成的容器镜像地址
IMAGE_ADDR=${IMAGE_HOST}/${NAMESPACE}/${SHORTNAME}:${IMAGE_VERSION}

function clean() {
  rm -rf deploy/systemd/
  rm -rf deploy/supervisor/
  rm -rf deploy/kubernetes/${DEPLOY_ENV}
}

function systemd() {
  mkdir -p deploy/systemd/
  cp -rf scripts/templates/systemd/microservice.service deploy/systemd/${APPNAME}.service

  eval "$SED" "s#_SERVICE_CODE_#${SERVICE_CODE}#g" deploy/systemd/${APPNAME}.service
  eval "$SED" "s#_PRODUCT_CODE_#${PRODUCT_CODE}#g" deploy/systemd/${APPNAME}.service
  eval "$SED" "s#_SHORT_NAME_#${SHORT_NAME}#g" deploy/systemd/${APPNAME}.service
  eval "$SED" "s#_API_VERSION_#${API_VERSION}#g" deploy/systemd/${APPNAME}.service
  eval "$SED" "s#_APPNAME_#${APPNAME}#g" deploy/systemd/${APPNAME}.service
}

function supervisor() {
  mkdir -p deploy/supervisor/
  cp -rf scripts/templates/supervisor/microservice.conf deploy/supervisor/${APPNAME}.conf

  eval "$SED" "s#_SERVICE_CODE_#${SERVICE_CODE}#g" deploy/supervisor/${APPNAME}.conf
  eval "$SED" "s#_PRODUCT_CODE_#${PRODUCT_CODE}#g" deploy/supervisor/${APPNAME}.conf
  eval "$SED" "s#_SHORT_NAME_#${SHORT_NAME}#g" deploy/supervisor/${APPNAME}.conf
  eval "$SED" "s#_API_VERSION_#${API_VERSION}#g" deploy/supervisor/${APPNAME}.conf
  eval "$SED" "s#_APPNAME_#${APPNAME}#g" deploy/supervisor/${APPNAME}.conf
}

function kubernetes() {
  mkdir -p deploy/kubernetes/${DEPLOY_ENV}/
  mkdir -p deploy/kubernetes/${DEPLOY_ENV}/config/configmap/
  cp -rf scripts/templates/kubernetes/* deploy/kubernetes/${DEPLOY_ENV}/

  if test -f config/app-${DEPLOY_ENV}-${BUILD_ENV}.yaml; then
    cp -a config/app-${DEPLOY_ENV}-${BUILD_ENV}.yaml deploy/kubernetes/${DEPLOY_ENV}/config/configmap/app.yaml
  fi

  eval "${SED}" "s#DEPLOY_ENV#${DEPLOY_ENV}#g" deploy/kubernetes/${DEPLOY_ENV}/kustomization.yaml
  eval "${SED}" "s#NAMESPACE#${NAMESPACE}#g" deploy/kubernetes/${DEPLOY_ENV}/kustomization.yaml
  eval "${SED}" "s#NAMESPACE#${NAMESPACE}#g" deploy/kubernetes/${DEPLOY_ENV}/workloads/deployment.yaml
  eval "${SED}" "s#NAMESPACE#${NAMESPACE}#g" deploy/kubernetes/${DEPLOY_ENV}/service/ingresses.yaml
  eval "${SED}" "s#DEPLOY_ENV#${DEPLOY_ENV}#g" deploy/kubernetes/${DEPLOY_ENV}/service/ingresses.yaml
  eval "${SED}" "s#IMAGE_NAME#${IMAGE_NAME}#g" deploy/kubernetes/${DEPLOY_ENV}/kustomization.yaml
  eval "${SED}" "s#IMAGE_NAME#${IMAGE_NAME}#g" deploy/kubernetes/${DEPLOY_ENV}/workloads/deployment.yaml
  eval "${SED}" "s#IMAGE_VERSION#${IMAGE_VERSION}#g" deploy/kubernetes/${DEPLOY_ENV}/kustomization.yaml
}

clean
systemd
supervisor
kubernetes
