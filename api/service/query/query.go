package query

import (
	"fmt"
	"github.com/dchest/uniuri"
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

func NewUUID() string {
	return uniuri.NewLen(uniuri.UUIDLen)
}

// Constants //
const (
	AUTH_TOKEN_DURATION = time.Hour
)

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

func (q Query) HandleUnique(handle string) bool {
	return !q.UserExistsByHandle(handle)
}

func (q Query) EmailUnique(email string) bool {
	return !q.UserExistsByEmail(email)
}

func (q Query) UserExistsByHandle(handle string) bool {
	found := []struct {
		Handle string `json:"u.handle"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: `
            MATCH   (u:User)
            WHERE   u.handle = {handle}
            RETURN  u.handle
        `,
		Parameters: neoism.Props{
			"handle": handle,
		},
		Result: &found,
	})
	return len(found) > 0
}

func (q Query) UserExistsByEmail(email string) bool {
	found := []struct {
		Email string `json:"u.email"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: `
            MATCH   (u:User)
            WHERE   u.email = {email}
            RETURN  u.email
        `,
		Parameters: neoism.Props{
			"email": email,
		},
		Result: &found,
	})
	return len(found) > 0
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

func (q Query) GetHashedPassword(handle string) (hashedPassword []byte, ok bool) {
	found := []struct {
		HashedPassword string `json:"u.password"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: `
            MATCH   (u:User)
            WHERE   u.handle = {handle}
            RETURN  u.password
        `,
		Parameters: neoism.Props{
			"handle": handle,
		},
		Result: &found,
	})

	if ok := len(found) > 0; !ok {
		return []byte{}, ok
	} else {
		return []byte(found[0].HashedPassword), ok
	}
}

func (q Query) SetGetNewAuthTokenForUser(handle string) (string, bool) {
	created := []struct {
		Token string `json:"a.value"`
	}{}
	now := Now()
	token := "Token " + NewUUID()
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: `
                MATCH   (u:User)
                WHERE   u.handle     = {handle}
                WITH    u
                OPTIONAL MATCH (u)<-[old_r:SESSION_OF]-(old_a:AuthToken)
                DELETE  old_r, old_a
                WITH    u
                CREATE  (u)<-[r:SESSION_OF]-(a:AuthToken)
                SET     r.created_at = {now}
                SET     a.value      = {token}
                SET     a.expires    = {time}
                RETURN  a.value
            `,
		Parameters: neoism.Props{
			"handle": handle,
			"token":  token,
			"time":   now.Add(AUTH_TOKEN_DURATION),
			"now":    now,
		},
		Result: &created,
	})
	if ok := len(created) > 0; ok {
		return created[0].Token, ok
	} else {
		return "", ok
	}
}
