FROM golang:1.23-bookworm AS builder
WORKDIR /app
COPY go.mod ./
# go.sum will be regenerated on first `go mod tidy` since we ship only go.mod.
RUN go mod download || true
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -o /out/server ./cmd/server

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /out/server /app/server
COPY migrations /app/migrations
EXPOSE 3001
CMD ["/app/server"]
