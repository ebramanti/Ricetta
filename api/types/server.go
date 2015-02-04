package types

import (
	"github.com/jadengore/Ricetta/api/service"
	"github.com/mccoyst/validate"
)

//
// Constructor Types for API Server
//

type Api struct {
	Svc       *service.Svc
	Validator *validate.V
}
