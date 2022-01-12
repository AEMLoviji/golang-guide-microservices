package app

import (
	"github.com/aemloviji/golang-guide-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
