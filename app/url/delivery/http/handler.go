package urlHttp

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
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
	switch true {
	case errors.Is(err, urlInterfaces.ErrAlreadyExists):
		w.WriteHeader(http.StatusConflict)
		return
	case err == nil:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetPage(w http.ResponseWriter, r *http.Request) {
	newUrl := mux.Vars(r)["url"]

	url, err := h.useCase.GetUrl(newUrl)
	switch true {
	case errors.Is(err, urlInterfaces.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	case err == nil:
		http.Redirect(w, r, url.OldUrl, http.StatusMovedPermanently)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
