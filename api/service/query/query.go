package query

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"time"
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

func (q Query) cypherOrPanic(query *neoism.CypherQuery) {
	panicIfErr(q.Db.Cypher(query))
}

func Now() time.Time {
	return time.Now().Local()
}

// Initializes the Neo4j Database
func (q Query) DatabaseInit() {
	if curator := q.CreatePublicCurator(); curator == nil {
		fmt.Println("Curator Node not initialized")
	}
}

func (q Query) CreatePublicCurator() *neoism.Node {
	if curator, _, err := q.Db.GetOrCreateNode("PublicCurator", "name", neoism.Props{
		"name": "PublicCurator",
	}); err != nil {
		panic(err)
	} else {
		panicIfErr(curator.AddLabel("PublicCurator"))
		return curator
	}
}

func (q Query) CreateUser(handle, email, passwordHash string) bool {
	newUser := []struct {
		Handle string    `json:"u.handle"`
		Email  string    `json:"u.email"`
		Joined time.Time `json:"u.joined"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: `
            CREATE (u:User {
                handle:   {handle},
                name:     "",
                email:    {email},
                password: {password},
                joined:   {joined}
            })
            RETURN u.handle, u.email, u.joined
        `,
		Parameters: neoism.Props{
			"handle":   handle,
			"email":    email,
			"password": passwordHash,
			"joined":   Now(),
		},
		Result: &newUser,
	})
	return len(newUser) > 0
}
