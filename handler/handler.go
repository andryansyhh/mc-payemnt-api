package handler

import (
	"mc-payment-api/usecase"
)

type handlerHttpServer struct {
	app usecase.UsecaseInterface
}

func NewHttpHandler(app usecase.UsecaseInterface) handlerHttpServer {
	return handlerHttpServer{app: app}
}
