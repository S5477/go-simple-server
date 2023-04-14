package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
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

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
