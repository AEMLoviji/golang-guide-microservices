package services

import (
	"github.com/aemloviji/golang-guide-microservices/mvc/domain"
	"github.com/aemloviji/golang-guide-microservices/mvc/utils"
)

type usersService struct{}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)

	//user, err := domain.UserDao.GetUser(userId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return user, nil
}
