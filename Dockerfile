FROM registry.cn-hangzhou.aliyuncs.com/opsaid/minideb:stretch

WORKDIR /opt

COPY build/service /opt/service
COPY config/app-mini.yaml /opt/config/app.yaml

EXPOSE 10080/tcp
EXPOSE 10081/tcp

ENTRYPOINT [ "/opt/service" ]
CMD [ "--config", "/opt/config/app.yaml" ]
