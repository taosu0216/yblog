FROM golang:1.23 AS builder

COPY . /src
COPY _posts/ /src/_posts/
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update -y && apt upgrade -y && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase util-linux \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/_posts /src/_posts
COPY --from=builder /src/dist /src/dist
COPY --from=builder /src/blacklist /src/blacklist

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./blug", "-conf", "/data/conf/config.yaml"]
