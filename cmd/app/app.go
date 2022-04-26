package app

import (
	envDelivery "github.com/e4t4g/Backend_Go_2lvl/internal/pkg/environment/delivery"
	envRepo "github.com/e4t4g/Backend_Go_2lvl/internal/pkg/environment/repository/sqlite"
	envUsecase "github.com/e4t4g/Backend_Go_2lvl/internal/pkg/environment/usecase"
	userDelivery "github.com/e4t4g/Backend_Go_2lvl/internal/pkg/user/delivery"
	userRepo "github.com/e4t4g/Backend_Go_2lvl/internal/pkg/user/repository/sqlite"
	userUsecase "github.com/e4t4g/Backend_Go_2lvl/internal/pkg/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

func App() {
	router := gin.Default()

	router.Run()

	db, err := sqlx.Open("sqlite3", "/DB.sql")
	if err != nil {
		log.Fatalf("can not connect to DB: %s", err)
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			log.Fatalf("can not connect to DB: %s", err)
		}
	}(db)

	userRepos := userRepo.New(db)
	envRepos := envRepo.New(db)

	userUsecases := userUsecase.New(userRepos)
	envUsecases := envUsecase.New(envRepos)

	usersDelivery := userDelivery.New(userUsecases)
	envsDelivery := envDelivery.New(envUsecases)

	router.POST("/", usersDelivery.Create())
	router.POST("/env", envsDelivery.Create())
	router.GET("/:id", usersDelivery.Find())
	router.GET("/env/:id", envsDelivery.Find())
	router.GET("/:id", usersDelivery.Delete())
	router.GET("/env/:id", envsDelivery.Delete())
	router.GET("/:id", usersDelivery.Add())
	router.GET("/env/:id", envsDelivery.Add())

}
