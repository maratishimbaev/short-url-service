package urlHttp

import (
	"net/http"
	"short-url-service/app/url/interfaces"
)

type Handler struct {
	useCase urlInterfaces.UrlUseCase
}

func NewHandler(useCase urlInterfaces.UrlUseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) AddUrl(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetPage(w http.ResponseWriter, r *http.Request) {

}
