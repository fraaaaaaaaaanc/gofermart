package main

import (
	"gofermart/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
