package usecase

import (
	"context"
	"github.com/e4t4g/Backend_Go_2lvl/internal/pkg/environment"
	"github.com/e4t4g/Backend_Go_2lvl/internal/pkg/models"
)

type usecase struct {
	repo environment.Repository
}

func New(repo environment.Repository) environment.Usecase {
	resp := usecase{repo: repo}
	return resp
}

func (u usecase) Create(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (u usecase) Add(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (u usecase) Delete(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (u usecase) Find(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

