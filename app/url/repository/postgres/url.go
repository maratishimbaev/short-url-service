package urlPostgres

import (
	"database/sql"
	"short-url-service/app/models"
	"short-url-service/app/url/interfaces"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (r *UrlRepository) CreateUrl(url *models.Url) (err error) {
	query := "insert into url (old_url, new_url) values ($1, $2)"
	_, err = r.db.Exec(query, &url.OldUrl, &url.NewUrl)
	if err != nil {
		return urlInterfaces.ErrAlreadyExists
	}

	return nil
}

func (r *UrlRepository) GetUrl(newUrl string) (url *models.Url, err error) {
	url = &models.Url{NewUrl: newUrl}

	query := "select old_url from url where new_url = $1"
	err = r.db.QueryRow(query, url.NewUrl).Scan(&url.OldUrl)
	if err != nil {
		return nil, urlInterfaces.ErrNotFound
	}

	return url, nil
}
