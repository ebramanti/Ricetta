package api

import (
	"github.com/jadengore/Ricetta/api/service"
	"github.com/jadengore/Ricetta/api/types"
	//"github.com/ant0ine/go-json-rest/rest"
)

type Api struct {
	Svc       *service.Svc
	Util      *util.Util
	Validator *validate.V
}

/**
 * Constructor
 */
func NewApi(uri string) *types.Api {
	api := &types.Api{
		Svc:       service.NewService(uri),
		Validator: types.NewValidator(),
	}
	return api
}
