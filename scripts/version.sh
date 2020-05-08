#!/bin/sh

source scripts/env

if test -z $1; then
    echo -e "Usage:"
    echo -e "\t ./scripts/version.sh prefix"
    echo -e "\t ./scripts/version.sh release"
    echo -e "\t ./scripts/version.sh update"
    exit 0;
fi

function prefix() {
    TEMP=$(grep "version: \".*\"" api/proto/${API_VERSION}/microservice.proto)
    PREFIX_VERSION=$(echo -n $TEMP | awk -F"\"" '{ print $2 }')
    echo $PREFIX_VERSION
}

function release() {
    TEMP=$(git describe --tags --dirty --always 2>/dev/null)
    RELEASE_VERSION=$TEMP

    if test -z $RELEASE_VERSION; then
        RELEASE_VERSION=v0.0.0
    fi

    echo $RELEASE_VERSION
}

function update() {
    GOHOSTS=$(go env GOHOSTOS)

    PREFIX_VERSION=$(prefix)
    RELEASE_VERSION=$(release)

    if test ${GOHOSTS} = "darwin"; then
        sed -i "" "s#version: \"${PREFIX_VERSION}\"#version: \"${RELEASE_VERSION}\"#g" api/proto/${API_VERSION}/microservice.proto
    else
        sed -i "s#version: \"${PREFIX_VERSION}\"#version: \"${RELEASE_VERSION}\"#g" api/proto/${API_VERSION}/microservice.proto
    fi
}

$1