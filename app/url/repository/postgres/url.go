package urlPostgres

import (
	"database/sql"
	"short-url-service/app/models"
	urlInterfaces "short-url-service/app/url/interfaces"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (r *UrlRepository) CreateUrl(url *models.Url) (err error) {
	query := "insert into url (old_url, new_url) values ($1, $2)"
	_, err = r.db.Exec(query, url.OldUrl, url.NewUrl)
	if err != nil {
		return urlInterfaces.ErrAlreadyExists
	}

	return nil
}

func (r *UrlRepository) GetUrl(oldUrl string) (url *models.Url, err error) {
	return nil, nil
}
