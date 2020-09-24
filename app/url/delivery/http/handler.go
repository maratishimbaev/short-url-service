package urlHttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"short-url-service/app/models"
	"short-url-service/app/url/interfaces"
)

type Handler struct {
	useCase urlInterfaces.UrlUseCase
}

func NewHandler(useCase urlInterfaces.UrlUseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) AddUrl(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var url models.Url

	err = json.Unmarshal(body, &url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.useCase.AddUrl(&url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetPage(w http.ResponseWriter, r *http.Request) {

}
