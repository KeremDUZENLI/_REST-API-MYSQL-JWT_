package main

import (
	"os"

	"jwt-project/common/env"

	"jwt-project/routes"
)

func main() {
	env.Load()

	port := os.Getenv("PORT")

	routes.Setup().Run(":" + port)
}
