package main

import (
	"github.com/cyamas/hodor-cache/client"
	"github.com/cyamas/hodor-cache/server"
)

func main() {
	go server.Start()
	client.Start()
}
