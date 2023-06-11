package utils

import (
	"playground/cmd/playground/underscore"
	_ "playground/cmd/playground/underscore"
)

func RunPkgWithInitFunc() {
	underscore.HelloWorld()
}
