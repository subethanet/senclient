package main

import (
	"manifold/con/listener"
	"fmt"
)


func main() {
	fmt.Println("Starting Manifold client.")
	server := listener.Create(4242)
	server.Start()
}