package types

import (
	"time"
)

//
// New User Type
//

type UserSignupProposal struct {
	Handle          string
	Email           string
	Password        string
	ConfirmPassword string
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
