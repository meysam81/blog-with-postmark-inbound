package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/meysam81/tarzan/cmd/metrics"
)

func (a *AppState) Websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := a.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		http.Error(w, "Failed upgrading the connection", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Connection close error:", err)
		}
	}()

	metrics.IncrementWebSocketConnections()

	if err := conn.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
		log.Println("Set read deadline error:", err)
		return
	}
	conn.SetPongHandler(func(string) error {
		if err := conn.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
			log.Println("Set read deadline error in pong handler:", err)
		}
		return nil
	})

	pingTicker := time.NewTicker(30 * time.Second)
	defer pingTicker.Stop()

	done := make(chan bool, 1)

	go func() {
		defer func() { done <- true }()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("WebSocket error: %v", err)
				}
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			metrics.DecrementWebSocketConnections()
			return
		case <-pingTicker.C:
			if err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
				log.Println("Set write deadline error:", err)
				return
			}
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping error:", err)
				return
			}
		case <-*a.Signal:
			log.Println("Notification received!")
			if err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
				log.Println("Set write deadline error:", err)
				return
			}
			err := conn.WriteMessage(websocket.TextMessage, []byte("New post available"))
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("Client disconnected:", err)
				} else {
					log.Println("Write error:", err)
				}
				return
			}
		}
	}
}
