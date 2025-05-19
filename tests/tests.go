package tests

import (
	"fmt"
	"server/bootstrap"
	"testing"

	"github.com/lucas11776-golang/http"
)

// TESTCASE HAS CHANGE...
// TODO Refactor HTTP to support application testing.
type Testing struct {
	HTTP *http.HTTP
}

func (ctx *Testing) Run() {
	fmt.Println("Running Application")
	ctx.HTTP.Listen()
}

func (ctx *Testing) Cleanup() {
	ctx.HTTP.Close()
	fmt.Println("Clean Up Done")
}

// TODO Looks like the structure meant work
func TestCase(t *testing.T) *Testing {
	testing := &Testing{HTTP: bootstrap.Boot(".env")}

	go testing.Run()

	t.Cleanup(testing.Cleanup)

	return testing
}
