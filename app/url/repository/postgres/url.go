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
	query := "insert into url (old_url, new_url) values ($1, $2)"
	_, err = r.db.Exec(query, url.OldUrl, url.NewUrl)
	if err != nil {
		return err
	}

	return nil
}

func (r *UrlRepository) GetUrl(oldUrl string) (url *models.Url, err error) {
	return nil, nil
}
