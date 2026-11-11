package main

import (
	"github.com/kamiertop/ktm/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
