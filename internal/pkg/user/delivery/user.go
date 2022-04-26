package delivery

import (
	"github.com/e4t4g/Backend_Go_2lvl/internal/pkg/user"
	"github.com/gin-gonic/gin"
)

type delivery struct {
	users user.Usecase
}

func New(users user.Usecase) user.Delivery {
	result := delivery{
		users: users,
	}
	return result
}

func (d delivery) Create() gin.HandlerFunc {
	panic("implement me")
}

func (d delivery) Add() gin.HandlerFunc {
	panic("implement me")
}

func (d delivery) Delete() gin.HandlerFunc {
	panic("implement me")
}

func (d delivery) Find() gin.HandlerFunc {
	panic("implement me")
}
