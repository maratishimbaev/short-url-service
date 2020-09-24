package urlInterfaces

import "short-url-service/app/models"

type UrlUseCase interface {
	AddUrl(url *models.Url) (err error)
	GetUrl(newUrl string) (url *models.Url, err error)
}
