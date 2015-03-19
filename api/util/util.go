package util

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type Util struct{}

func (u Util) GetTokenFromHeader(r *rest.Request) string {
	return r.Header.Get("Authorization")
}
