//Server-A application
//
//  the purpose of this application is to provide a test server
//  for traefik forwardauth
//
//     Schemes: http, https
//     BasePath: /api/v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
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
