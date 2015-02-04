package query

import (
	"fmt"
	"github.com/jmcvetta/neoism"
)

// Query is a private type, and stored locally to package.

type Query struct {
	Db *neoism.Database
}

func NewQuery(uri string) *Query {
	neo4jdb, err := neoism.Connect(uri)
	panicIfErr(err)

	query := Query{neo4jdb}
	query.DatabaseInit()

	return &query
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Initializes the Neo4j Database
func (q Query) DatabaseInit() {
	fmt.Println("Initialize database here")
}
