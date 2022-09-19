package main

import (
	"mypackage/app/controllers"
)

func main() {
	// controllers.MemosController()
	// controllers.MeigensController()
	router := controllers.GetRouter()
	router.Run()

	// net/httpを使用する場合
	// http.HandleFunc("/tasks_net_http", controllers.GetAllTasksNetHttp)
	// http.ListenAndServe(":8080", nil)

}
