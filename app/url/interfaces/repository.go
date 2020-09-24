package urlInterfaces

import "short-url-service/app/models"

type UrlRepository interface {
	CreateUrl(url *models.Url) (err error)
	GetUrl(newUrl string) (url *models.Url, err error)
}
