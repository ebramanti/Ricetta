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
