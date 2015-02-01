package api

import (
	"github.com/jadengore/Ricetta/api/types"
	//"github.com/ant0ine/go-json-rest/rest"
)

/**
 * Constructor
 */
func NewApi(uri string) *types.Api {
	api := &types.Api{
	//Svc:       service.NewService(uri),
	//Validator: types.NewValidator(),
	}
	return api
}
