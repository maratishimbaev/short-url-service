package usecase

import (
	"math/rand"
	"short-url-service/app/models"
	"short-url-service/app/url/interfaces"
)

type UrlUseCase struct {
	repository urlInterfaces.UrlRepository
}

func NewUrlUseCase(repository urlInterfaces.UrlRepository) *UrlUseCase {
	return &UrlUseCase{repository: repository}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateUrl(size uint64) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

func (u *UrlUseCase) AddUrl(url *models.Url) (err error) {
	if url.NewUrl == "" {
		url.NewUrl = generateUrl(10)
	}

	return u.repository.CreateUrl(url)
}

func (u *UrlUseCase) GetUrl(oldUrl string) (url *models.Url, err error) {
	return nil, nil
}
