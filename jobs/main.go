package main

import (
	"server/jobs/office"
)

type Jobs struct {
	Office *office.Office
}

func main() {
	office.NewOffice().Launch()

	select {}
}
