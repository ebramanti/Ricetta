var recipe = require('./Recipe/recipe')

module.exports = function (server) {

	server.route({
		method: 'POST',
		path: '/recipes',
		handler: recipe.NewRecipe,
		config: {
			validate: {
				payload: recipe.Schema
			}
		}
	})

	server.route({
		method: 'GET',
		path: '/recipes',
		// handler: function (req, reply) {
		// 	reply('hello!')
		// },
		handler: recipe.GetRecipes
		// config: {}
	})

}
