package main

import (
	"fmt"

	"github.com/niwek/kwl-forwardauth/config"
	"github.com/niwek/kwl-forwardauth/controller"
	"github.com/niwek/kwl-forwardauth/log"
)

func main() {
	routerEngine := controller.Router.Setup()

	if err := routerEngine.Run(":" + config.Env.Port); err != nil {
		log.Fatal(
			fmt.Sprintf("Error starting service on port %s: %s", config.Env.Port, err.Error()),
		)
	}
}
