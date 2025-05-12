# worker/Dockerfile
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./

FROM gcr.io/distroless/base
COPY --from=builder /app/worker /worker
ENTRYPOINT ["/worker"]
