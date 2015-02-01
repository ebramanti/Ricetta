package routes

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jadengore/Ricetta/api/types"
)

func MakeHandler(api types.Api, disableLogs bool) (rest.ResourceHandler, error) {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		DisableLogger:            disableLogs,
	}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}},
	)

	return handler, err
}
