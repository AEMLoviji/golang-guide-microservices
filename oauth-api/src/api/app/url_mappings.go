package app

import (
	"github.com/aemloviji/golang-guide-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
}
