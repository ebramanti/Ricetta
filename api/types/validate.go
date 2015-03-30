package types

import (
	"github.com/jadengore/goconfig"
	"github.com/mccoyst/validate"
	"regexp"
)

type vc struct {
	// Config Constants
	MAX_HANDLE_LENGTH       int
	MIN_PASS_LENGTH         int
	MAX_PASS_LENGTH         int
	MAX_RECIPE_TITLE_LENGTH int
	MAX_RECIPE_NOTES_LENGTH int
	AUTH_TOKEN_EXPIRES      int64

	// Regex Constants
	HANDLE_REGEX    *regexp.Regexp
	EMAIL_REGEX     *regexp.Regexp
	TIME_UNIT_REGEX *regexp.Regexp
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
		"handle":      vd.validateHandle,
		"email":       vd.validateEmail,
		"password":    vd.validatePassword,
		"timeunit":    vd.validateTimeUnit,
		"time":        vd.validateTime,
		"recipetitle": vd.validateRecipeTitle,
		"recipenotes": vd.validateRecipeNotes,
	}
	return &vd
}

func initializeConstants(config *goconfig.ConfigFile) vc {
	c := vc{}
	// Constants set by configuration
	c.MAX_HANDLE_LENGTH, _ = config.GetInt("global", "handle-length")
	c.MIN_PASS_LENGTH, _ = config.GetInt("global", "min-pass")
	c.MAX_PASS_LENGTH, _ = config.GetInt("global", "max-pass")
	c.MAX_RECIPE_TITLE_LENGTH, _ = config.GetInt("global", "max-recipe-title")
	c.MAX_RECIPE_NOTES_LENGTH, _ = config.GetInt("global", "max-recipe-notes")

	// Using int64 for query layer
	c.AUTH_TOKEN_EXPIRES, _ = config.GetInt64("global", "auth-token-expires")

	// Regular expression constants
	c.HANDLE_REGEX = regexp.MustCompile(`^[\p{L}\p{M}][\d\p{L}\p{M}]*$`)
	c.EMAIL_REGEX = regexp.MustCompile(`^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	c.TIME_UNIT_REGEX = regexp.MustCompile(`(hr|min|sec|day|week)s?$`)
	return c
}
