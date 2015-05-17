var db = require('../connector')
var uuid = require('uuid').v4

var LABEL = 'Recipe'

module.exports = function (resource) {

  resource.NewRecipe = function (req, reply) {
    var recipe = req.payload
    recipe.uuid = uuid()
    db.save(recipe, function (err, node) {
      if (err) { return reply(err, null); }

      db.label(node, [ LABEL ], function (err) {
        if (err) { return reply(err, null); }

        return reply(recipe)
          .code(201)
      })
    })
  }

  resource.GetRecipes = function (req, reply) {
    db.nodesWithLabel(LABEL, function (err, results) {
      return reply(results)
        .code(200)
    })
  }

}

