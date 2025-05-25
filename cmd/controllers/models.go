package controllers

import (
	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/datastore"
)

type AppState struct {
	Config *config.Config
	DS     datastore.Datastore
}
