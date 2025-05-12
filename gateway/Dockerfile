# gateway/Dockerfile
# 1. Build stage
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# produce a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux go build -o gateway ./ 

# 2. Runtime stage
FROM gcr.io/distroless/base
COPY --from=builder /app/gateway /gateway
ENTRYPOINT ["/gateway"]
