package api

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type Api struct {
	Svc       *service.Svc
	Util      *apiutil.Util
	Validator *validate.V
}

/**
 * Constructor
 */
func NewApi(uri string) *Api {
	api := &Api{
		Svc:       service.NewService(uri),
		Util:      &apiutil.Util{},
		Validator: types.NewValidator(),
	}
	return api
}
