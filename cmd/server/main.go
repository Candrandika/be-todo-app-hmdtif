package main

import (
	"log"

	"github.com/Candrandika/be-todo-app-hmdtif/bootstrap"
)

func main() {
	app := bootstrap.App()
	log.Fatal(app.Listen(":" + app.Env.AppPort))
}
