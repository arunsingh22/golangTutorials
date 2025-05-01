package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func serverWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v\n", err)
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	log.Println("New WebSocket connection established")

	for {
		// Close handler to detect when the client disconnects
		conn.SetCloseHandler(func(code int, text string) error {
			log.Printf("Client disconnected: Code=%d, Reason=%s\n", code, text)
			return nil
		})

		if err := conn.WriteMessage(ws.PingMessage, nil); err != nil {
			fmt.Println("Client disconnected, hence stopping response", err)
			break
		}
		// Echo message back to client
		t := time.Now().String()
		// response := fmt.Sprintf("Current system time %s", message)
		err = conn.WriteMessage(ws.TextMessage, []byte(t))
		if err != nil {
			log.Printf("Error writing message: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	// log.Println("WebSocket connection closed")
}

func main() {
	http.HandleFunc("/ws", serverWS)
	log.Println("WebSocket server started on :9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalf("Server exited with error: %v\n", err)
	}
}
