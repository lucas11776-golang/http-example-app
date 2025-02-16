package main

import (
	"fmt"
	"server/bootstrap"
)

func main() {
	server := bootstrap.Boot()

	fmt.Printf("Server Running %s", server.Host())

	server.Listen()
}
