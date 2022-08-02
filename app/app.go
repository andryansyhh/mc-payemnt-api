package app

import (
	"fmt"
	"mc-payment-api/infra"
	postgres "mc-payment-api/infra/postgre"
	"mc-payment-api/repository"
	"mc-payment-api/usecase"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	Repo := repository.NewRepo(postgres.PSQL.DB.DB)
	app := usecase.NewUsecase(Repo)
	infra.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
