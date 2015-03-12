package types

import (
	"github.com/jadengore/goconfig"
	"github.com/mccoyst/validate"
	"regexp"
)

type vc struct {
	// Config Constants
	MAX_HANDLE_LENGTH int
	MIN_PASS_LENGTH   int
	MAX_PASS_LENGTH   int

	// Regex Constants
	HANDLE_REGEX *regexp.Regexp
	EMAIL_REGEX  *regexp.Regexp
}

// Main Ricetta Validator
// Validator function found in respective data
// package in types

type RicettaValidator struct {
	Validator validate.V
	Constants vc
}

func NewValidator(config *goconfig.ConfigFile) *RicettaValidator {
	vd := RicettaValidator{}
	vd.Constants = initializeConstants(config)
	vd.Validator = validate.V{
		"handle":   vd.validateHandle,
		"email":    vd.validateEmail,
		"password": vd.validatePassword,
	}
	return &vd
}

func initializeConstants(config *goconfig.ConfigFile) vc {
	c := vc{}
	// Constants set by configuration
	c.MAX_HANDLE_LENGTH, _ = config.GetInt("global", "handle-length")
	c.MIN_PASS_LENGTH, _ = config.GetInt("global", "min-pass")
	c.MAX_PASS_LENGTH, _ = config.GetInt("global", "max-pass")

	// Regular expression constants
	c.HANDLE_REGEX = regexp.MustCompile(`^[\p{L}\p{M}][\d\p{L}\p{M}]*$`)
	c.EMAIL_REGEX = regexp.MustCompile(`^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	c.TIME_UNIT_REGEX = regexp.MustCompile(`(hr|min|sec|day|week)s?$`)
	return c
}
