package urlPostgres

import (
	"database/sql"
	"short-url-service/app/models"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (r *UrlRepository) CreateUrl(url *models.Url) (err error) {
	return nil
}

func (r *UrlRepository) GetUrl(oldUrl string) (url *models.Url, err error) {
	return nil, nil
}
