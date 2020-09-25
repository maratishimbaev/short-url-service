package urlHttp

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"short-url-service/app/middleware"
	"short-url-service/app/models"
	urlInterfaces "short-url-service/app/url/interfaces"
	"short-url-service/app/url/usecase/mock"
	"testing"
)

type urlSuite struct {
	suite.Suite
	router *mux.Router
	controller *gomock.Controller
	useCase *mock.MockUrlUseCase
	url models.Url
	urlByte *bytes.Buffer
	urlWithoutNew models.Url
	urlWithoutNewByte *bytes.Buffer
}

func (suite *urlSuite) SetupTest() {
	suite.router = mux.NewRouter()
	suite.router.Use(middleware.RecoveryMiddleware)
	suite.router.Use(middleware.LogMiddleware)

	suite.controller = gomock.NewController(suite.T())
	suite.useCase = mock.NewMockUrlUseCase(suite.controller)

	RegisterHttpEndpoints(suite.router, suite.useCase)

	suite.url = models.Url{
		OldUrl: "https://google.com",
		NewUrl: "goo",
	}
	urlJson, err := json.Marshal(suite.url)
	suite.urlByte = bytes.NewBuffer(urlJson)
	assert.NoError(suite.T(), err)

	suite.urlWithoutNew = models.Url{
		OldUrl: "https://google.com",
	}
	urlWithoutNewJson, err := json.Marshal(suite.urlWithoutNew)
	suite.urlWithoutNewByte = bytes.NewBuffer(urlWithoutNewJson)
	assert.NoError(suite.T(), err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(urlSuite))
}

func (suite *urlSuite) TestAddUrlSuccess() {
	suite.useCase.EXPECT().
		AddUrl(gomock.Any()).
		Return(nil).
		Times(1)

	r, _ := http.NewRequest("POST", "/", suite.urlByte)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *urlSuite) TestAddUrlWithoutNewSuccess() {
	suite.useCase.EXPECT().
		AddUrl(gomock.Any()).
		Return(nil).
		Times(1)

	r, _ := http.NewRequest("POST", "/", suite.urlWithoutNewByte)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *urlSuite) TestAddUrlConflict() {
	suite.useCase.EXPECT().
		AddUrl(gomock.Any()).
		Return(urlInterfaces.ErrAlreadyExists).
		Times(1)

	r, _ := http.NewRequest("POST", "/", suite.urlByte)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	assert.Equal(suite.T(), http.StatusConflict, w.Code)
}

func (suite *urlSuite) TestAddUrlEmpty() {
	suite.useCase.EXPECT().
		AddUrl(gomock.Any()).
		Return(nil).
		Times(1)

	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *urlSuite) TestGetPageSuccess() {
	suite.useCase.EXPECT().
		GetUrl(gomock.Any()).
		Return(&models.Url{}, nil).
		Times(1)

	r, _ := http.NewRequest("GET", "/goo", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	assert.Equal(suite.T(), http.StatusMovedPermanently, w.Code)
}

func (suite *urlSuite) TestGetPageNotFound() {
	suite.useCase.EXPECT().
		GetUrl(gomock.Any()).
		Return(nil, urlInterfaces.ErrNotFound).
		Times(1)

	r, _ := http.NewRequest("GET", "/gooo", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}
