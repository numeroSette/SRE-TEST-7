package register

import (
	getrandomnumber "github.com/numeroSette/SRE-TEST-7/cmd/get-random-number"
	"github.com/numeroSette/SRE-TEST-7/internal/router"
)

func init() {
	router.Router.HandleFunc("/random-number", getrandomnumber.GetRandomNumber)
}
