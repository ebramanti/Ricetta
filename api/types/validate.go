package types

import (
	"github.com/mccoyst/validate"
)

func NewValidator() *validate.V {
	vd := validate.V{}
	return &vd
}
