package app

import (
	"github.com/aemloviji/golang-guide-microservices/oauth-api/src/api/controllers/oauth"
	"github.com/aemloviji/golang-guide-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
