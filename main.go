package main

import (
	"seredaes/go-passmem/App/DB"
	"seredaes/go-passmem/App/Env"
	"seredaes/go-passmem/App/Web"
)

func main() {

	// LOAD ENV
	Env.LoadEnv()

	// RESTORE DB INTO MEMORY
	DB.RestoreDB()

	// START WEB
	serverAddress, _ := Env.GetConfig("SERVER")
	Web.StartWebServer(serverAddress)

}
