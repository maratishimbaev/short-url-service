package server

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	_ "github.com/lib/pq"
	"math/rand"
	"net/http"
	"os"
	"short-url-service/app/middleware"
	urlHttp "short-url-service/app/url/delivery/http"
	"short-url-service/app/url/interfaces"
	"short-url-service/app/url/repository/postgres"
	"short-url-service/app/url/usecase"
	"time"
)

type app struct {
	urlUseCase urlInterfaces.UrlUseCase
}

func InitDatabase() (db *sql.DB, err error) {
	dbInfo := fmt.Sprintf(
		"host=127.0.0.1 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("SHORT_URL_SERVICE_USER"),
		os.Getenv("SHORT_URL_SERVICE_PASSWORD"),
		os.Getenv("SHORT_URL_SERVICE_DBNAME"),
	)

	return sql.Open("postgres", dbInfo)
}

func NewApp() *app {
	db, err := InitDatabase()
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	err = db.Ping()
	if err != nil {
		golog.Error(err.Error())
		return nil
	}

	urlRepository := urlPostgres.NewUrlRepository(db)

	return &app{
		urlUseCase: usecase.NewUrlUseCase(urlRepository),
	}
}

var port = flag.Uint64("p", 8000, "port")

func (a *app) Start() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	router := mux.NewRouter()
	router.Use(middleware.LogMiddleware)

	urlHttp.RegisterHttpEndpoints(router, a.urlUseCase)

	http.Handle("/", router)

	golog.Infof("Server started at port %d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		golog.Error("Server failed: ", err.Error())
	}
}
