package app

import (
	"github.com/aemloviji/golang-guide-microservices/src/api/controllers/polo"
	"github.com/aemloviji/golang-guide-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}
