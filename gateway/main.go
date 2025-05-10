package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/google/uuid"
    "github.com/nats-io/nats.go"
)

// Order represents the payload we'll send
type Order struct {
    ID        string    `json:"id"`
    Item      string    `json:"item"`
    CreatedAt time.Time `json:"created_at"`
}

func main() {
    // 1. Connect to NATS
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatalf("NATS connect failed: %v", err)
    }
    defer nc.Drain()

    // 2. HTTP handler
    http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "POST only", http.StatusMethodNotAllowed)
            return
        }

        // parse item from JSON body
        var payload struct{ Item string `json:"item"` }
        if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
            http.Error(w, "bad payload", http.StatusBadRequest)
            return
        }

        // build an Order with a traceable ID
        ord := Order{
            ID:        uuid.NewString(),
            Item:      payload.Item,
            CreatedAt: time.Now().UTC(),
        }

        data, _ := json.Marshal(ord)
        if err := nc.Publish("orders", data); err != nil {
            log.Printf("publish error: %v", err)
            http.Error(w, "publish failed", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    })

    log.Println("gateway listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
