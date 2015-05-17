var handler = {}
var db = require('../connector')
var uuid = require('uuid').v4

handler.NewRecipe = function (req, reply) {
  var recipe = req.payload
  recipe.uuid = uuid()
  db.save(recipe, function (err, node) {
    db.label(node, ['Recipe'], function (err) {
      if (err) {
        reply('Failed to create recipe', null)
          .code(500)
      } else {
        reply(recipe)
          .code(201)
      }
    })
  })
}

module.exports = handler