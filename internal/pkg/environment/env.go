package environment

import (
	"context"
	"github.com/e4t4g/Backend_Go_2lvl/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

type Delivery interface {
	Create() gin.HandlerFunc
	Add() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Find() gin.HandlerFunc
}

type Usecase interface {
	Create(ctx context.Context, user *models.Environment) error
	Add(ctx context.Context, user *models.Environment) error
	Delete(ctx context.Context, user *models.Environment) error
	Find(ctx context.Context, user *models.Environment) error
}

type Repository interface {
	Create(ctx context.Context, user *models.Environment) error
	Add(ctx context.Context, user *models.Environment) error
	Delete(ctx context.Context, user *models.Environment) error
	Find(ctx context.Context, user *models.Environment) error
}
