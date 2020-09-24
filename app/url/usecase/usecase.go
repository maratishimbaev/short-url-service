package usecase

import (
	"short-url-service/app/models"
	"short-url-service/app/url/interfaces"
)

type UrlUseCase struct {
	repository urlInterfaces.UrlRepository
}

func NewUrlUseCase(repository urlInterfaces.UrlRepository) *UrlUseCase {
	return &UrlUseCase{repository: repository}
}

func (u *UrlUseCase) AddUrl(newUrl string) (err error) {
	return nil
}

func (u *UrlUseCase) GetUrl(oldUrl string) (url *models.Url, err error) {
	return nil, nil
}
