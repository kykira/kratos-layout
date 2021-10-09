FROM golang:1.17 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM alpine:latest

COPY --from=builder /src/bin /data
COPY ./configs /data/conf
WORKDIR /data

#EXPOSE 8000
EXPOSE 9000
#VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
