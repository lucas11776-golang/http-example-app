package main

import (
	"fmt"
	"server/bootstrap"
	"server/env"
)

func main() {
	env.Load(".env")

	server := bootstrap.Boot()

	fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	server.Listen()
}
