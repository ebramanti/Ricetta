package types

import (
	"time"
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
