package register

import (
	getrandomnumber "github.com/testing/sre-test-1/cmd/get-random-number"
	"github.com/testing/sre-test-1/internal/router"
)

func init() {
	router.Router.HandleFunc("/random-number", getrandomnumber.GetRandomNumber)
}
