package types

import (
	"github.com/mccoyst/validate"
)

func NewValidator() *validate.V {
	vd := validate.V{}
type vc struct {
	// Config Constants
	MAX_HANDLE_LENGTH int
	MIN_PASS_LENGTH   int
	MAX_PASS_LENGTH   int

	// Regex Constants
	HANDLE_REGEX *regexp.Regexp
	EMAIL_REGEX  *regexp.Regexp
}

type RicettaValidator struct {
	Validator validate.V
	Constants vc
}
	return &vd
}
