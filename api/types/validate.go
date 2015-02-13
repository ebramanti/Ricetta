package types

import (
	"fmt"
	"github.com/jadengore/goconfig"
	"github.com/mccoyst/validate"
	"regexp"
	"unicode/utf8"
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

type RicettaValidator struct {
	Validator validate.V
	Constants vc
}

func NewValidator(config *goconfig.ConfigFile) *RicettaValidator {
	vd := RicettaValidator{}
	vd.Constants = initializeConstants(config)
	vd.Validator = validate.V{
		"handle": vd.validateHandle,
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
	return c
}

func (v RicettaValidator) validateHandle(i interface{}) error {
	handle := i.(string)
	if handle == "" {
		return fmt.Errorf("Required field for signup")
	} else if utf8.RuneCountInString(handle) > v.Constants.MAX_HANDLE_LENGTH {
		return fmt.Errorf("Too long, max length is %d", v.Constants.MAX_HANDLE_LENGTH)
	} else if !v.Constants.HANDLE_REGEX.MatchString(handle) {
		return fmt.Errorf(handle + " contains illegal characters")
	} else {
		return nil
	}
}
