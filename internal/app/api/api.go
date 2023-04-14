package api

import "github.com/sirupsen/logrus"

type API struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

func (api *API) Run() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("starting api server at port:", api.config.BindAddr)

	return nil
}
