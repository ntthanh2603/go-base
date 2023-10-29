package main

import (
	"go-base/bootstrap"
	"go-base/configs"
)

func main() {
	app := bootstrap.App()

	app.Run(configs.Port)
}
