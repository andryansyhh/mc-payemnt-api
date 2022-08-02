package main

import (
	"mc-payment-api/app"
	infra "mc-payment-api/infra/postgre"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	if err := infra.InitPostgre(); err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
