FROM alpine:3.11

RUN apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata

COPY dist/api /data/console/api

CMD /data/console/api
