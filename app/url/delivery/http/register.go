package urlHttp

import (
	"github.com/gorilla/mux"
	"short-url-service/app/url/interfaces"
)

func RegisterHttpEndpoints(router *mux.Router, useCase urlInterfaces.UrlUseCase) {
	h := NewHandler(useCase)

	router.HandleFunc("/", h.AddUrl).Methods("POST")
	router.HandleFunc("/{url:[a-zA-Z0-9]}", h.GetPage).Methods("GET")
}
