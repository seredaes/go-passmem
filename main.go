package main

import (
	"seredaes/go-passmem/App/Env"
	Logger "seredaes/go-passmem/App/Log"
)

func main() {

	// Load ENV
	Env.LoadEnv()

	Logger.Log("Test", "danger")
	Logger.Log("Test", "danger")

}
