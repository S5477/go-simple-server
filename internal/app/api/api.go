package api

import (
	"net/http"

	"github.com/S5477/go-simple-server/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Run() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.configureRouterField()
	if err := api.configureStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
