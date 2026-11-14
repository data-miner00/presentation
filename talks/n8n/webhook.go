package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	ticket, _ := json.Marshal(map[string]string{
		"from":    "alex@example.com",
		"subject": "Checkout page returns 500",
		"body":    "Nobody has been able to pay since this morning.",
	})
	url := "http://127.0.0.1:5678/webhook-test/support-ticket"

	resp, err := http.Post(url, "application/json", bytes.NewReader(ticket))
	if err != nil {
		fmt.Println("Is n8n running?", err)
		return
	}
	defer resp.Body.Close()
	reply, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.Status, string(reply))
}
