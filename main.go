package main

import (
	"github.com/dispatchlabs/disgo/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
