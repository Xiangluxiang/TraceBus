## Run & Test Locally

1. Start NATS

```sh
docker run -d --rm -p 4222:4222 nats
```

2. Run worker

```sh
cd ~/projects/tracebus/worker
go run main.go
```

3. Run gateway

```sh
cd ~/projects/tracebus/gateway
go run main.go
```

4. Send a test order

```sh
curl -XPOST localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"item":"latte"}'
```
