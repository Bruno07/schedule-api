package main

import (

	"github.com/Bruno07/schedule/internal/application/http/routes"
	"github.com/Bruno07/schedule/internal/config"
)

func main() {

	config.LoadConfig()

	routes.LoadRoutes().Run(config.GetPort())

}
