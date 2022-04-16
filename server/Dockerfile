FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/gin-vue-admin
COPY . .
RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/gin-vue-admin
COPY --from=0 /go/src/gin-vue-admin/server ./
COPY --from=0 /go/src/gin-vue-admin/config.yaml ./
COPY --from=0 /go/src/gin-vue-admin/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server

apiVersion: v1
kind: ReplicationController
metadata:
    name: sqltools-ycz
    labels:
    name: sqltools-ycz
spec:
    replicas: 1
    selector:
    name: sqltools-ycz
    template:
    metadata:
        labels:
        name: sqltools-ycz
    spec:
        containers:
        - name: sqltools-ycz
        image: myhub.fdccloud.com/common/sqltools
        ports:
        - containerPort: 80
        env:
        - name: DB_HOST
            value: "mysql-ycz.default.svc.cluster.local"
        - name: DB_NAME
            value: "config"
        - name: DB_USERNAME
            value: "ycz"
        - name: DB_PASSWORD
            value: "77TqrfPaMbwaZXYu"
        - name: CACHE_HOST
            value: "memcached-ycz.default.svc.cluster.local"
        - name: CACHE_PORT
            value: "11211"
        - name: CACHE_PREFIX
            value: "ycz"
        nodeSelector:
        kubernetes.io/hostname: k8s-master
        imagePullSecrets:
        - name: registrykey-myhub
