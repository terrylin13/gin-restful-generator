package ws

import (
	"log"
	"net/http"
	"time"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	userID := r.URL.Query().Get("user_id")
	connManager.Lock()
	client := &Client{
		Conn: conn,
		Msg:  make(chan []byte),
	}
	connManager.connections[userID] = client
	connManager.Unlock()

	go func() {
		defer func() {
			conn.Close()
			connManager.Lock()
			delete(connManager.connections, userID)
			connManager.Unlock()
		}()

		// Sending heartbeat messages
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := client.Send([]byte("ping")); err != nil {
					log.Println("Heartbeat failed: ", err)
					return
				}
			case msg := <-client.Msg:
				if err := client.Send(msg); err != nil {
					log.Println("Send msg failed: ", err)
					return
				}
			}
		}
	}()

	// Reading incoming messages
	go func() {
		defer func() {
			conn.Close()
			connManager.Lock()
			delete(connManager.connections, userID)
			connManager.Unlock()
		}()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read failed:", err)
				break
			}
		}
	}()
}
