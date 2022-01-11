package services

import (
	"net/http"

	"github.com/aemloviji/golang-guide-microservices/mvc/domain"
	"github.com/aemloviji/golang-guide-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemId int64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
