package main

import (
	"Cobain/config"
	"Cobain/routes"
)

func main() {
	config.Init()
	config.InitDB()
	e := routes.Routing()
	e.Logger.Info(e.Start(":8000"))
}
