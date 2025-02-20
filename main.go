package main

import (
	"fmt"
	"server/bootstrap"
	query "server/query_test"
)

func main() {
	server := bootstrap.Boot()

	query.Test()

	fmt.Printf("Server Running %s", server.Host())

	server.Listen()
}
