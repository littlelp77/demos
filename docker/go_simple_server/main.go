package main

import (
	"github.com/gin-contrib/cors"
	"github.com/simple-server/routes"
)

func main() {

	router := routes.Routes()
	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	router.Use(cors.New(config))
	//router.Use(cors.Default())
	router.Run(":8085")
}
