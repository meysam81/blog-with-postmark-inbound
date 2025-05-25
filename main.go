package main

import (
	"embed"

	"github.com/meysam81/tarzan/cmd/server"
)

//go:embed dist/*
var frontend embed.FS

func main() {
	server.Main(frontend)
}
