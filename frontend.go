package main

import (
	"embed"
)

//go:embed dist/*
var frontend embed.FS
