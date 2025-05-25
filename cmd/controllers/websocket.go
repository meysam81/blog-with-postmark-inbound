package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (a *AppState) Websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := a.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		http.Error(w, "Failed upgrading the connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		select {
		case _ = <-*a.Signal:
			log.Println("Notification received!")
			err := conn.WriteMessage(websocket.TextMessage, []byte("New post available"))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}
}
