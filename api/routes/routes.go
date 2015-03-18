package routes

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jadengore/Ricetta/api/api"
)

func MakeHandler(api api.Api, disableLogs bool) (rest.ResourceHandler, error) {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		DisableLogger:            disableLogs,
	}

	err := handler.SetRoutes(
		&rest.Route{"POST", "/signup", api.Signup},
		&rest.Route{"POST", "/login", api.Login},
	)

	return handler, err
}
