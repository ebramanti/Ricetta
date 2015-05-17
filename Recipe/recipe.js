var recipe = {}

recipe.Schema = require('./schema')

require('./handler')(recipe)

module.exports = recipe