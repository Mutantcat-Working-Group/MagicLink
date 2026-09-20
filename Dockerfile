# MagicLink 容器镜像：多阶段构建，产出 CGO_ENABLED=0 静态二进制。
# 该工具通过符号链接补全指令，运行时需挂载包含 /mlink 的目录，例如：
#   docker run --rm -v "$PWD":/work -w /work ghcr.io/mutantcat-working-group/magiclink:latest magiclink --help
FROM golang:1.24-alpine AS build

ARG VERSION=1.0.20260920

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.version=${VERSION}" \
    -o /out/magiclink .

FROM alpine:3.21

RUN adduser -D -H -u 10001 magiclink \
    && mkdir -p /work \
    && chown magiclink:magiclink /work

COPY --from=build /out/magiclink /usr/local/bin/magiclink

USER magiclink
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/magiclink"]
