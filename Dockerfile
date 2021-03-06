FROM docker:dind
WORKDIR /
EXPOSE 4200
RUN mkdir -p /etc/docker
RUN > /etc/docker/daemon.json
RUN echo "{\"registry-mirrors\":[\"https://mirror.ccs.tencentyun.com\",\"https://hub-mirror.c.163.com\"]}" > /etc/docker/daemon.json
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add docker-compose
ADD . .