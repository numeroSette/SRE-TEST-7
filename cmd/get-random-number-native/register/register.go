package register

import (
	getrandomnumbernative "github.com/numeroSette/SRE-TEST-7/cmd/get-random-number-native"
	"github.com/numeroSette/SRE-TEST-7/internal/router"
)

func init() {
	router.Router.HandleFunc("/get-random-number-native", getrandomnumbernative.GetRandomNumberNative)
}
