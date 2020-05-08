FROM {{IMAGE_FROM}}

WORKDIR /opt

COPY build/service /opt/service
COPY config/app-sample.toml /opt/config/app.toml

EXPOSE 10080/tcp
EXPOSE 10081/tcp

ENTRYPOINT [ "/opt/service" ]
CMD [ "--config", "/opt/config/app.toml" ]