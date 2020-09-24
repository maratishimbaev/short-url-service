package urlInterfaces

import "short-url-service/app/models"

type UrlUseCase interface {
	AddUrl(newUrl string) (err error)
	GetUrl(oldUrl string) (url *models.Url, err error)
}
