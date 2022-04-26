package sqlite

import (
	"context"
	"github.com/e4t4g/Backend_Go_2lvl/internal/pkg/environment"
	"github.com/e4t4g/Backend_Go_2lvl/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}


func New(db *sqlx.DB) environment.Repository{
	resp := repository{db: db}

	return resp
}

func (r repository) Create(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (r repository) Add(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (r repository) Find(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}
