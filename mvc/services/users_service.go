package services

import (
	"github.com/aemloviji/golang-guide-microservices/mvc/domain"
	"github.com/aemloviji/golang-guide-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
