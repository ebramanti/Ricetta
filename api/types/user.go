package types

import (
	"types"
)

//
// New User Type
//

type UserSignup struct {
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
