# Build context is the repo root: docker build -f services/orders/Dockerfile .

# Stage 1: Build
FROM golang:1.26.6 AS builder
WORKDIR /usr/src/app

COPY pkg ./pkg
COPY services/orders ./services/orders
# trimmed workspace: shared pkg + this service only
RUN go work init ./pkg ./services/orders
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o main ./services/orders/cmd

# Stage 2: Production
FROM gcr.io/distroless/static-debian12:nonroot AS deploy
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/main ./main
EXPOSE 8082
CMD ["./main"]
