package util

import (
	"github.com/jadengore/Ricetta/api/types"
)

func (u Util) MakeRecipeUrl(recipeid string) string {
	return "http://ricetta.io/recipes/" + recipeid
}

func (u Util) AddRecipeUrl(r types.Recipe) types.Recipe {
	r.URL = u.MakeRecipeUrl(r.Id)
	return r
}

func (u Util) AddRecipeUrlToArray(recipes types.Recipes) types.Recipes {
	for index, recipe := range recipes {
		updated := u.AddRecipeUrl(recipe)
		recipes[index] = updated
	}
	return recipes
}
