package api

type API struct {
	config *Config
}

func New(config *Config) *API {
	return &API{
		config: config,
	}
}

func (api *API) Run() error {
	return nil
}