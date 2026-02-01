# syntax=docker/dockerfile:1.6
FROM golang:1.25-alpine AS builder

ENV GOPROXY=https://proxy.golang.org,direct

WORKDIR /app
COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=$(go env GOARCH) \
    go build \
    -ldflags="-w -s -extldflags '-static'" \
    -trimpath \
    -buildvcs=false \
    -o bot ./cmd/bot
FROM alpine:3.18

RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow

WORKDIR /app
COPY --from=builder /app/bot .
COPY --from=builder /app/migrations/ ./migrations/
COPY --from=builder /app/script/ ./script/

EXPOSE 8118
CMD ["./bot"]