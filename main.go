package main

import (
	"go-demo/api"
)

func main() {
	r := api.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
