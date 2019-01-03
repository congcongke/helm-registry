FROM cargo.congcongke.io/congcongke/debian:jessie

RUN echo "Asia/Shanghai" > /etc/timezone

RUN mkdir /data

COPY bin/registry /registry

COPY image/config.yaml /config.yaml

CMD ["/registry","serve","-c","/config.yaml"]
