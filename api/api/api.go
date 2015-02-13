package api

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jadengore/Ricetta/api/service"
	"github.com/jadengore/Ricetta/api/types"
	"github.com/jadengore/Ricetta/api/util"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Api struct {
	Svc       *service.Svc
	Util      *util.Util
	Validator *types.RicettaValidator
}

/**
 * Constructor
 */
func NewApi(uri string) *Api {
	api := &Api{
		Svc:       service.NewService(uri),
		Util:      &util.Util{},
		Validator: types.NewValidator(),
	}
	return api
}

//
// Begin API functions
//

func (a Api) Signup(w rest.ResponseWriter, r *rest.Request) {
	proposal := types.UserSignupProposal{}
	if err := r.DecodeJsonPayload(&proposal); err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if err := a.Validator.ValidateAndTag(proposal, "json"); err != nil {
	// 	// Return error in specific field
	// }

	handle := proposal.Handle
	email := proposal.Email
	password := proposal.Password
	confirm_password := proposal.ConfirmPassword

	if password != confirm_password {
		a.Util.SimpleJsonReason(w, 403, "Passwords do not match")
		return
	}

	// TODO: Implement unique handle/email

	// Ensure unique handle
	// if unique := a.Svc.HandleIsUnique(handle); !unique {
	// 	a.Util.SimpleJsonReason(w, 409, "Sorry, handle or email is already taken")
	// 	return
	// }

	// // Ensure unique email
	// if unique := a.Svc.EmailIsUnique(email); !unique {
	// 	a.Util.SimpleJsonReason(w, 409, "Sorry, handle or email is already taken")
	// 	return
	// }

	var hashed_pass string

	if hash, err := bcrypt.GenerateFromPassword([]byte(password), 10); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		hashed_pass = string(hash)
	}

	if !a.Svc.CreateNewUser(handle, email, hashed_pass) {
		a.Util.SimpleJsonReason(w, http.StatusInternalServerError, "Unexpected failure to create new user")
		return
	}

	w.WriteHeader(201)
	w.WriteJson(types.Json{
		"response": "Signed up a new user!",
		"handle":   handle,
		"email":    email,
	})

}
