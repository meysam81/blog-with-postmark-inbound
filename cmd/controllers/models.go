package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/datastore"
	"github.com/meysam81/tarzan/cmd/filestore"
)

type AppState struct {
	Config         *config.Config
	DS             datastore.Datastore
	Upgrader       websocket.Upgrader
	Signal         *chan uint8
	Filestore      filestore.Filestore
	AttachmentPath string
}
