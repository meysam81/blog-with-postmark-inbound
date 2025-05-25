package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/datastore"
)

type AppState struct {
	Config   *config.Config
	DS       datastore.Datastore
	Upgrader websocket.Upgrader
	Signal   *chan uint8
}
