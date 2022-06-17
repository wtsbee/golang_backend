package main

import (
	"mypackage/app/controllers"
)

func main() {
	// controllers.MemosController()
	// controllers.MeigensController()
	router := controllers.GetRouter()
	router.Run()
}
