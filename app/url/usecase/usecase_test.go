package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"short-url-service/app/models"
	"short-url-service/app/url/repository/mock"
	"testing"
)

type urlSuite struct {
	suite.Suite
	controller *gomock.Controller
	useCase *UrlUseCase
	repository *mock.MockUrlRepository
	url models.Url
	urlWithoutNew models.Url
}

func (suite *urlSuite) SetupTest() {
	suite.controller = gomock.NewController(suite.T())
	defer suite.controller.Finish()

	suite.repository = mock.NewMockUrlRepository(suite.controller)
	suite.useCase = NewUrlUseCase(suite.repository)

	suite.url = models.Url{
		OldUrl: "https://google.com",
		NewUrl: "goo",
	}

	suite.urlWithoutNew = models.Url{
		OldUrl: "https://google.com",
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(urlSuite))
}

func (suite *urlSuite) TestAddUrl() {
	suite.repository.EXPECT().CreateUrl(&suite.url).Return(nil).Times(1)

	err := suite.useCase.AddUrl(&suite.url)

	assert.NoError(suite.T(), err)
}

func (suite *urlSuite) TestAddUrlWithoutNew() {
	suite.repository.EXPECT().CreateUrl(&suite.urlWithoutNew).Return(nil).Times(1)

	err := suite.useCase.AddUrl(&suite.urlWithoutNew)

	assert.NoError(suite.T(), err)
}

func (suite *urlSuite) TestGetUrl() {
	suite.repository.EXPECT().GetUrl(suite.url.NewUrl).Return(&suite.url, nil).Times(1)

	_, err := suite.useCase.GetUrl(suite.url.NewUrl)

	assert.NoError(suite.T(), err)
}
