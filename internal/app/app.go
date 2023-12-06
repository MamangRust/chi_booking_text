package app

import (
	"booking_chi_text/internal/handler"
	"booking_chi_text/internal/repository"
	"booking_chi_text/internal/service"
	"booking_chi_text/pkg/hash"
	"net/http"
)

func Run() {
	hashing := hash.NewHashingPassword()

	repository := repository.NewRepository()

	service := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashing,
	})

	myhandler := handler.NewHandler(service)

	r := myhandler.Init()

	http.ListenAndServe(":8080", r)

}
