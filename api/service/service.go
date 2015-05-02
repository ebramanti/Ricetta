package service

import (
	"github.com/jadengore/Ricetta/api/service/query"
	"github.com/jadengore/Ricetta/api/types"
	"github.com/jadengore/Ricetta/api/util"
	"github.com/jadengore/goconfig"
)

type Svc struct {
	Query *query.Query
	Util  *util.Util
}

func NewService(uri string, config *goconfig.ConfigFile) *Svc {
	s := &Svc{
		query.NewQuery(uri, config),
		&util.Util{},
	}
	return s
}

func (s Svc) CreateNewUser(handle, email, passwordHash string) bool {
	return s.Query.CreateUser(handle, email, passwordHash)
}

func (s Svc) HandleIsUnique(handle string) bool {
	return s.Query.HandleUnique(handle)
}

func (s Svc) EmailIsUnique(email string) bool {
	return s.Query.EmailUnique(email)
}

func (s Svc) GetHashedPassword(handle string) (hashedPassword []byte, ok bool) {
	return s.Query.GetHashedPassword(handle)
}

func (s Svc) SetGetNewAuthToken(handle string) (token string, ok bool) {
	return s.Query.SetGetNewAuthTokenForUser(handle)
}

func (s Svc) DestroyAuthToken(token string) bool {
	return s.Query.DestroyAuthToken(token)
}

func (s Svc) VerifyAuthToken(token string) bool {
	return s.Query.FindAuthToken(token)
}

func (s Svc) UserExists(handle string) bool {
	return s.Query.UserExistsByHandle(handle)
}

func (s Svc) GetHandleFromAuthorization(token string) (handle string, ok bool) {
	return s.Query.DeriveHandleFromAuthToken(token)
}

func (s Svc) NewRecipe(handle string, recipe types.Recipe) (res types.Recipe, ok bool) {
	if recipe, ok := s.Query.CreateRecipe(handle, recipe); ok {
		recipe = s.Util.AddRecipeUrl(recipe)
		return recipe, ok
	} else {
		return types.Recipe{}, ok
	}
}

func (s Svc) GetOwnRecipes(handle string) (recipes types.Recipes, ok bool) {
	if ok := s.UserExists(handle); ok {
		if recipes, ok := s.Query.GetOwnRecipes(handle); ok {
			recipes = s.Util.AddRecipeUrlToArray(recipes)
			return recipes, ok
		} else {
			return types.Recipes{}, ok
		}

	} else {
		return types.Recipes{}, ok
	}
}

func (s Svc) GetCuratedRecipes() (recipes types.Recipes, ok bool) {
	if recipes, ok := s.Query.GetCuratedRecipes(); ok {
		recipes = s.Util.AddRecipeUrlToArray(recipes)
		return recipes, ok
	} else {
		return types.Recipes{}, ok
	}
}

func (s Svc) GetVisibleRecipeById(handle, id string) (recipe types.Recipe, ok bool) {
	if recipe, ok := s.Query.GetVisibleRecipeById(handle, id); ok {
		recipe = s.Util.AddRecipeUrl(recipe)
		return recipe, ok
	} else {
		return types.Recipe{}, ok
	}
}

func (s Svc) GetCuratedRecipeById(id string) (recipe types.Recipe, ok bool) {
	if recipe, ok := s.Query.GetCuratedRecipeById(id); ok {
		recipe = s.Util.AddRecipeUrl(recipe)
		return recipe, ok
	} else {
		return types.Recipe{}, ok
	}
}
