package types

import (
	"fmt"
	"time"
	"unicode/utf8"
)

//
// New User Type
//

type UserSignupProposal struct {
	Handle          string `json:"handle" validate:"handle"`
	Email           string `json:"email" validate:"email"`
	Password        string `json:"password" validate:"password"`
	ConfirmPassword string `json:"confirmpassword" validate:"password"`
}

type InitialUserAttributes struct {
	FirstName string
	LastName  string
	Gender    string
	Bio       string
	Birthday  time.Time
}

//
// Account Types
//

type UserLogin struct {
	Handle   string
	Password string
}

//
// Validator Functions
//

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

func (v RicettaValidator) validateEmail(i interface{}) error {
	email := i.(string)
	if email == "" {
		return fmt.Errorf("Required field for signup")
	} else if !v.Constants.EMAIL_REGEX.MatchString(email) {
		return fmt.Errorf(email + " is an invalid email")
	} else {
		return nil
	}
}

func (v RicettaValidator) validatePassword(i interface{}) error {
	password := i.(string)
	passwordLen := utf8.RuneCountInString(password)
	if password == "" {
		return fmt.Errorf("Required field for signup")
	} else if passwordLen < v.Constants.MIN_PASS_LENGTH {
		return fmt.Errorf("Too short, minimum length is %d", v.Constants.MIN_PASS_LENGTH)
	} else if passwordLen > v.Constants.MAX_PASS_LENGTH {
		return fmt.Errorf("Too long, maximum length is %d", v.Constants.MAX_PASS_LENGTH)
	} else {
		return nil
	}
}
