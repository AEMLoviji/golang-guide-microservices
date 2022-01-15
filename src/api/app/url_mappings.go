package app

import (
	"github.com/aemloviji/golang-guide-microservices/src/api/controllers/polo"
	"github.com/aemloviji/golang-guide-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
