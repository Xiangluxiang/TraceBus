package main

import (
    "encoding/json"
    "log"

    "github.com/nats-io/nats.go"
)

type Order struct {
    ID        string `json:"id"`
    Item      string `json:"item"`
    CreatedAt string `json:"created_at"`
}

func main() {
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatalf("NATS connect failed: %v", err)
    }
    defer nc.Drain()

    // subscribe to the "orders" subject
    nc.Subscribe("orders", func(msg *nats.Msg) {
        var ord Order
        if err := json.Unmarshal(msg.Data, &ord); err != nil {
            log.Printf("unmarshal error: %v", err)
            return
        }

        log.Printf("[worker] received order %s for item %q (created %s)",
            ord.ID, ord.Item, ord.CreatedAt)

        // TODO: simulate work, emit downstream event, etc.
    })
    nc.Flush()

    log.Println("worker listening for orders…")
    select {} // block forever
}
