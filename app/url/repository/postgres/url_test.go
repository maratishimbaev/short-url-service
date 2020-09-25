package urlPostgres

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"short-url-service/app/models"
	urlInterfaces "short-url-service/app/url/interfaces"
	"testing"
)

type urlSuite struct {
	suite.Suite
	repository *UrlRepository
	db *sql.DB
	mock sqlmock.Sqlmock
	url models.Url
}

func (suite *urlSuite) SetupTest() {
	var err error
	suite.db, suite.mock, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.repository = NewUrlRepository(suite.db)

	suite.url = models.Url{
		OldUrl: "https://google.com",
		NewUrl: "goo",
	}
}

func (suite *urlSuite) TearDown() {
	assert.NoError(suite.T(), suite.db.Close())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(urlSuite))
}

func (suite *urlSuite) TestCreateUrlSuccess() {
	suite.mock.
		ExpectExec("insert into url").
		WithArgs(suite.url.OldUrl, suite.url.NewUrl).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repository.CreateUrl(&suite.url)

	assert.NoError(suite.T(), err)
}

func (suite *urlSuite) TestCreateUrlAlreadyExists() {
	suite.mock.
		ExpectExec("insert into url").
		WithArgs(suite.url.OldUrl, suite.url.NewUrl).
		WillReturnError(urlInterfaces.ErrAlreadyExists)

	err := suite.repository.CreateUrl(&suite.url)

	assert.Error(suite.T(), err)
}

func (suite *urlSuite) TestGetUrlSuccess() {
	rows := sqlmock.NewRows([]string{"old_url"}).AddRow(suite.url.OldUrl)

	suite.mock.
		ExpectQuery("select old_url from url").
		WithArgs(suite.url.NewUrl).
		WillReturnRows(rows)

	_, err := suite.repository.GetUrl(suite.url.NewUrl)

	assert.NoError(suite.T(), err)
}

func (suite *urlSuite) TestGetUrlNotFound() {
	suite.mock.
		ExpectQuery("select old_url from url").
		WithArgs(suite.url.NewUrl).
		WillReturnError(urlInterfaces.ErrNotFound)

	_, err := suite.repository.GetUrl("not_found_url")

	assert.Error(suite.T(), err)
}
