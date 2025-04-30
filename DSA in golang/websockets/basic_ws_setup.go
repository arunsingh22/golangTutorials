package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin:     func(r *http.Request) bool { return true },
// }

// func serverWS(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Printf("WebSocket upgrade error: %v\n", err)
// 		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close()

// 	log.Println("New WebSocket connection established")

// 	for {
// 		// Read message from client
// 		// mt, message, err := conn.ReadMessage()
// 		// if err != nil {
// 		// 	log.Printf("Error reading message: %v\n", err)
// 		// 	break
// 		// }
// 		// log.Printf("Received: %s\n", string(message))

// 		// Echo message back to client
// 		t := time.Now().String()
// 		// response := fmt.Sprintf("Current system time %s", message)
// 		err = conn.WriteMessage(websocket.TextMessage, []byte(t))
// 		if err != nil {
// 			log.Printf("Error writing message: %v\n", err)
// 		}
// 		time.Sleep(2 * time.Second)
// 	}

// 	// log.Println("WebSocket connection closed")
// }

// func main() {
// 	http.HandleFunc("/ws", serverWS)
// 	log.Println("WebSocket server started on :9000")

// 	err := http.ListenAndServe(":9000", nil)
// 	if err != nil {
// 		log.Fatalf("Server exited with error: %v\n", err)
// 	}
// }
