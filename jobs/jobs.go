package jobs

import (
	"context"
	"server/jobs/office"
)

type Jobs struct {
	Office *office.Office
}

func Setup() {
	office.NewOffice().Launch(context.Background())

	select {}
}
