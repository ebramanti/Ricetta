var handler = {}
var db = require('../connector')
var uuid = require('uuid').v4

handler.NewRecipe = function (req, reply) {
  var recipe = req.payload
  recipe.uuid = uuid()
  db.save(recipe, function (err, node) {
    if (err) { return reply(err, null); }

    db.label(node, ['Recipe'], function (err) {
      if (err) { return reply(err, null); }
      
      return reply(recipe)
        .code(201)
    })
  })
}

module.exports = handler