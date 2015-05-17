var recipe = require('./Recipe/recipe')

var routes = module.exports = [
	{
		method: 'POST',
		path: '/recipes',
		handler: recipe.NewRecipe,
		config: {
			validate: {
				payload: recipe.Schema
			}
		}
	}
]
