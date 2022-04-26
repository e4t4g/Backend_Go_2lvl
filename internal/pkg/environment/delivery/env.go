package delivery

import (
	"github.com/e4t4g/Backend_Go_2lvl/internal/pkg/environment"
	"github.com/gin-gonic/gin"
)

type delivery struct {
	users environment.Usecase
}

func New(env environment.Usecase) environment.Delivery {
	resp := delivery{users: env}
	return resp
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

