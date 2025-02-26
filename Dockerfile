FROM golang:1.17 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn,direct
RUN go build go-fly.go

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /app/config

CMD ["/app/go-fly","server"]