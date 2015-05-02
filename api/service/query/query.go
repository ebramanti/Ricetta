package query

import (
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/jadengore/Ricetta/api/types"
	"github.com/jadengore/goconfig"
	"github.com/jmcvetta/neoism"
	"io/ioutil"
	"time"
)

type QueryStrings struct {
	FindToken         string
	CreateRecipe      string
	ReturnRecipe      string
	FindCurator       string
	AddCuratorRel     string
	CreateIngredient  string
	CreateStep        string
	GetOwnRecipes     string
	GetIngredients    string
	GetSteps          string
	GetCuratedRecipes string
	GetOwnRecipe      string
	GetVisibleRecipe  string
	RecipeAuthor      string
	TagExists         string
	CreateNewTag      string
	AddRelationToTag  string
}

// Query is a private type, and stored locally to package.

type Query struct {
	Db *neoism.Database
	Vd *types.RicettaValidator
	Qs QueryStrings
}

const (
	NANOSECONDS_IN_DAY int64 = 86400000000000
	CQL_DIR                  = "service/query/cql/"
)

var (
	expires time.Duration
)

func NewQuery(uri string, config *goconfig.ConfigFile) *Query {
	neo4jdb, err := neoism.Connect(uri)
	panicIfErr(err)

	query := Query{
		neo4jdb,
		types.NewValidator(config),
		QueryStringInit(),
	}

	query.DatabaseInit()
	query.ConstantInit()
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

func QueryStringInit() QueryStrings {
	return QueryStrings{
		FindToken:         parseQueryString(CQL_DIR + "findtoken.cql"),
		CreateRecipe:      parseQueryString(CQL_DIR + "createrecipenode.cql"),
		ReturnRecipe:      parseQueryString(CQL_DIR + "returnrecipenode.cql"),
		FindCurator:       parseQueryString(CQL_DIR + "findcurator.cql"),
		AddCuratorRel:     parseQueryString(CQL_DIR + "addcuratorrel.cql"),
		CreateIngredient:  parseQueryString(CQL_DIR + "createingredientnode.cql"),
		CreateStep:        parseQueryString(CQL_DIR + "createstepnode.cql"),
		GetOwnRecipes:     parseQueryString(CQL_DIR + "getownrecipes.cql"),
		GetCuratedRecipes: parseQueryString(CQL_DIR + "getcuratedrecipes.cql"),
		GetIngredients:    parseQueryString(CQL_DIR + "getingredients.cql"),
		GetSteps:          parseQueryString(CQL_DIR + "getsteps.cql"),
		GetOwnRecipe:      parseQueryString(CQL_DIR + "getownrecipe.cql"),
		GetVisibleRecipe:  parseQueryString(CQL_DIR + "getvisiblerecipe.cql"),
		RecipeAuthor:      parseQueryString(CQL_DIR + "recipeauthor.cql"),
		TagExists:         parseQueryString(CQL_DIR + "tagexists.cql"),
		CreateNewTag:      parseQueryString(CQL_DIR + "createnewtag.cql"),
		AddRelationToTag:  parseQueryString(CQL_DIR + "addrelationtotag.cql"),
	}
}

func parseQueryString(filename string) string {
	queryString, err := ioutil.ReadFile(filename)
	panicIfErr(err)
	return string(queryString)
}

// Initializes the Neo4j Database
func (q Query) DatabaseInit() {
	if curator := q.CreatePublicCurator(); curator == nil {
		fmt.Println("Curator Node not initialized")
	}
}

// Initializes constants needed for query layer
func (q Query) ConstantInit() {
	expires = time.Duration(q.Vd.Constants.AUTH_TOKEN_EXPIRES * NANOSECONDS_IN_DAY)
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
			"time":   now.Add(expires),
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

func (q Query) FindAuthToken(token string) bool {
	found := []struct {
		Handle string `json:"u.handle"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: q.Qs.FindToken,
		Parameters: neoism.Props{
			"token": token,
			"now":   Now(),
		},
		Result: &found,
	})
	return len(found) == 1
}

func (q Query) DeriveHandleFromAuthToken(token string) (handle string, ok bool) {
	found := []struct {
		Handle string `json:"u.handle"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: q.Qs.FindToken,
		Parameters: neoism.Props{
			"token": token,
			"now":   Now(),
		},
		Result: &found,
	})
	if ok = len(found) > 0; ok {
		return found[0].Handle, ok
	} else {
		return "", ok
	}
}

func (q Query) DestroyAuthToken(token string) bool {
	deleted := []struct {
		Handle string `json:"u.handle"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: `
            MATCH   (u:User)<-[so:SESSION_OF]-(a:AuthToken)
            WHERE   a.value = {token}
            DELETE  so, a
            RETURN  u.handle
        `,
		Parameters: neoism.Props{
			"token": token,
		},
		Result: &deleted,
	})
	return len(deleted) > 0
}

func (q Query) CreateRecipe(handle string, recipe types.Recipe) (res types.Recipe, ok bool) {
	recipeQuery := q.Qs.CreateRecipe
	if !recipe.Private {
		recipeQuery = q.Qs.FindCurator + recipeQuery + q.Qs.AddCuratorRel
	}
	createdRecipe := []types.Recipe{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: recipeQuery + q.Qs.ReturnRecipe,
		Parameters: neoism.Props{
			"handle":       handle,
			"id":           NewUUID(),
			"now":          Now(),
			"title":        recipe.Title,
			"notes":        recipe.Notes,
			"cooktime":     recipe.CookTime,
			"cooktimeunit": recipe.CookTimeUnit,
			"preptime":     recipe.PrepTime,
			"preptimeunit": recipe.PrepTimeUnit,
		},
		Result: &createdRecipe,
	})
	if ok = len(createdRecipe) > 0; !ok {
		return types.Recipe{}, ok
	} else {
		recipeid := createdRecipe[0].Id
		createdIngredients := make(types.Ingredients, len(recipe.Ingredients))
		for index, ingredient := range recipe.Ingredients {
			created := types.Ingredients{}
			q.cypherOrPanic(&neoism.CypherQuery{
				Statement: q.Qs.CreateIngredient,
				Parameters: neoism.Props{
					"rid":        recipeid,
					"id":         index + 1,
					"now":        Now(),
					"name":       ingredient.Name,
					"amount":     ingredient.Amount,
					"amountunit": ingredient.AmountUnit,
					"url":        ingredient.URL,
				},
				Result: &created,
			})
			if len(created) > 0 {
				createdIngredients[index] = created[0]
			}
		}
		if ok = (len(createdIngredients) == len(recipe.Ingredients)); !ok {
			return types.Recipe{}, !ok
		} else {
			createdSteps := make(types.Steps, len(recipe.Steps))
			for index, step := range recipe.Steps {
				created := types.Steps{}
				q.cypherOrPanic(&neoism.CypherQuery{
					Statement: q.Qs.CreateStep,
					Parameters: neoism.Props{
						"rid":         recipeid,
						"id":          index + 1,
						"now":         Now(),
						"instruction": step.Instruction,
						"time":        step.Time,
						"timeunit":    step.TimeUnit,
					},
					Result: &created,
				})
				if len(created) > 0 {
					createdSteps[index] = created[0]
				}
			}
			if ok = (len(createdSteps) == len(recipe.Steps)); !ok {
				return types.Recipe{}, !ok
			} else {
				createdTags := make(types.Tags, len(recipe.Tags))
				for index, tag := range recipe.Tags {
					created := types.Tags{}
					exists := types.Tags{}
					q.cypherOrPanic(&neoism.CypherQuery{
						Statement: q.Qs.TagExists,
						Parameters: neoism.Props{
							"name": tag.Name,
						},
						Result: &exists,
					})
					if len(exists) > 0 {
						q.cypherOrPanic(&neoism.CypherQuery{
							Statement: q.Qs.AddRelationToTag,
							Parameters: neoism.Props{
								"name": tag.Name,
								"rid":  recipeid,
							},
							Result: &created,
						})
					} else {
						q.cypherOrPanic(&neoism.CypherQuery{
							Statement: q.Qs.CreateNewTag,
							Parameters: neoism.Props{
								"name": tag.Name,
								"rid":  recipeid,
							},
							Result: &created,
						})
					}
					if len(created) > 0 {
						createdTags[index] = created[0]
					}
				}
				if ok = (len(createdTags) == len(recipe.Tags)); !ok {
					return types.Recipe{}, !ok
				} else {
					result := createdRecipe[0]
					result.Ingredients = createdIngredients
					result.Steps = createdSteps
					result.Tags = createdTags
					return result, ok
				}
			}
		}
	}
}

func (q Query) GetOwnRecipes(handle string) (res types.Recipes, ok bool) {
	recipes := types.Recipes{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: q.Qs.GetOwnRecipes,
		Parameters: neoism.Props{
			"handle": handle,
		},
		Result: &recipes,
	})
	for index, recipe := range recipes {
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetIngredients,
			Parameters: neoism.Props{
				"rid": recipe.Id,
			},
			Result: &recipes[index].Ingredients,
		})
	}
	for index, recipe := range recipes {
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetSteps,
			Parameters: neoism.Props{
				"rid": recipe.Id,
			},
			Result: &recipes[index].Steps,
		})
	}
	if ok := len(recipes) > 0; ok {
		return recipes, ok
	} else {
		return types.Recipes{}, ok
	}
}

func (q Query) GetCuratedRecipes() (res types.Recipes, ok bool) {
	recipes := types.Recipes{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: q.Qs.GetCuratedRecipes,
		Result:    &recipes,
	})
	for index, recipe := range recipes {
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetIngredients,
			Parameters: neoism.Props{
				"rid": recipe.Id,
			},
			Result: &recipes[index].Ingredients,
		})
	}
	for index, recipe := range recipes {
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetSteps,
			Parameters: neoism.Props{
				"rid": recipe.Id,
			},
			Result: &recipes[index].Steps,
		})
	}
	if ok := len(recipes) > 0; ok {
		return recipes, ok
	} else {
		return types.Recipes{}, ok
	}
}

func (q Query) GetVisibleRecipeById(handle string, id string) (r types.Recipe, ok bool) {
	check := []struct {
		Handle string `json:"handle"`
	}{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: q.Qs.RecipeAuthor,
		Parameters: neoism.Props{
			"handle": handle,
			"rid":    id,
		},
		Result: &check,
	})
	recipe := types.Recipes{}
	if check[0].Handle == handle {
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetOwnRecipe,
			Parameters: neoism.Props{
				"handle": handle,
				"rid":    id,
			},
			Result: &recipe,
		})
	} else {
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetVisibleRecipe,
			Parameters: neoism.Props{
				"rid": id,
			},
			Result: &recipe,
		})
	}
	if ok := len(recipe) > 0; ok {
		r = recipe[0]
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetIngredients,
			Parameters: neoism.Props{
				"rid": r.Id,
			},
			Result: &r.Ingredients,
		})
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetSteps,
			Parameters: neoism.Props{
				"rid": r.Id,
			},
			Result: &r.Steps,
		})
		return r, ok
	} else {
		return types.Recipe{}, ok
	}
}

func (q Query) GetCuratedRecipeById(id string) (r types.Recipe, ok bool) {
	recipe := types.Recipes{}
	q.cypherOrPanic(&neoism.CypherQuery{
		Statement: q.Qs.GetVisibleRecipe,
		Parameters: neoism.Props{
			"rid": id,
		},
		Result: &recipe,
	})
	if ok := len(recipe) > 0; ok {
		r = recipe[0]
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetIngredients,
			Parameters: neoism.Props{
				"rid": r.Id,
			},
			Result: &r.Ingredients,
		})
		q.cypherOrPanic(&neoism.CypherQuery{
			Statement: q.Qs.GetSteps,
			Parameters: neoism.Props{
				"rid": r.Id,
			},
			Result: &r.Steps,
		})
		return r, ok
	} else {
		return types.Recipe{}, ok
	}
}
