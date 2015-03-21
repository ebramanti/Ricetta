package api

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jadengore/Ricetta/api/service"
	"github.com/jadengore/Ricetta/api/types"
	"github.com/jadengore/Ricetta/api/util"
	"github.com/jadengore/goconfig"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Api struct {
	Svc  *service.Svc
	Util *util.Util
	Vd   *types.RicettaValidator
}

/**
 * Constructor
 */
func NewApi(uri string, config *goconfig.ConfigFile) *Api {
	api := &Api{
		Svc:  service.NewService(uri, config),
		Util: &util.Util{},
		Vd:   types.NewValidator(config),
	}
	return api
}

func (a Api) Authenticate(r *rest.Request) bool {
	if token := a.Util.GetTokenFromHeader(r); token != "" {
		return a.Svc.VerifyAuthToken(token)
	} else {
		return false
	}
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

	if err := a.Vd.Validator.ValidateAndTag(proposal, "json"); err != nil {
		a.Util.SimpleJsonValidationReason(w, 400, err)
		return
	}

	handle := proposal.Handle
	email := proposal.Email
	password := proposal.Password
	confirm_password := proposal.ConfirmPassword

	if password != confirm_password {
		a.Util.SimpleJsonReason(w, 403, "Passwords do not match")
		return
	}

	//Ensure unique handle
	if unique := a.Svc.HandleIsUnique(handle); !unique {
		a.Util.SimpleJsonReason(w, 409, "Sorry, handle or email is already taken")
		return
	}

	// Ensure unique email
	if unique := a.Svc.EmailIsUnique(email); !unique {
		a.Util.SimpleJsonReason(w, 409, "Sorry, handle or email is already taken")
		return
	}

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

func (a Api) Login(w rest.ResponseWriter, r *rest.Request) {
	credentials := types.UserLogin{}
	if err := r.DecodeJsonPayload(&credentials); err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.Vd.Validator.ValidateAndTag(credentials, "json"); err != nil {
		a.Util.SimpleJsonValidationReason(w, 400, err)
		return
	}

	handle := credentials.Handle
	password := []byte(credentials.Password)

	if passwordHash, ok := a.Svc.GetHashedPassword(handle); !ok {
		a.Util.SimpleJsonReason(w, 403, "Invalid username or password, please try again.")
		return
	} else {
		// err is nil if successful, error if comparison failed
		if err := bcrypt.CompareHashAndPassword(passwordHash, password); err != nil {
			a.Util.SimpleJsonReason(w, 403, "Invalid username or password, please try again.")
			return
		} else {
			// Create an Authentication token and return it to client
			if token, ok := a.Svc.SetGetNewAuthToken(handle); !ok {
				a.Util.SimpleJsonReason(w, 500, "Unexpected failure to produce new Authorization token")
			} else {
				w.WriteHeader(201)
				w.WriteJson(types.Json{
					"handle":   handle,
					"response": "User " + handle + " now logged in. Note your Authorization token.",
					"token":    token,
				})
				return
			}
		}
	}
}

func (a Api) Logout(w rest.ResponseWriter, r *rest.Request) {
	if a.Svc.DestroyAuthToken(a.Util.GetTokenFromHeader(r)) {
		w.WriteHeader(204)
		return
	} else {
		a.Util.SimpleJsonReason(w, 403, "Cannot invalidate token because it is missing")
		return
	}
}

func (a Api) NewRecipe(w rest.ResponseWriter, r *rest.Request) {

}
