package service

import (
	"github.com/jadengore/Ricetta/api/service/query"
)

type Svc struct {
	Query *query.Query
}

func NewService(uri string) *Svc {
	s := &Svc{
		query.NewQuery(uri),
	}
	return s
}

func (s Svc) CreateNewUser(handle, email, passwordHash string) bool {
	return s.Query.CreateUser(handle, email, passwordHash)
}
