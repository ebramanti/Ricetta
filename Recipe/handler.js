var uuid = require('uuid').v4
var async = require('async')

var db = require('../connector')
var step = require('./step')

var LABEL = 'Recipe'

module.exports = function (resource) {

  resource.NewRecipe = function (req, reply) {
    var recipe = req.payload
    var steps = recipe.steps
    recipe.uuid = uuid()
    recipe.steps = undefined
    async.waterfall([
      function createRecipeNode (callback) {
        db.save(recipe, function (err, node) {
          if (err) { return reply(err.message, null); }
          callback(null, node)
        })
      },
      function createRecipeLabel (node, callback) {
        db.label(node, [ LABEL ], function (err) {
          if (err) { return reply(err.message, null); }
          callback(null, node)
        })
      }
    ], function (error, node) {
        step.CreateSteps(node, steps, function(error) {
          if (error) {
            return reply('Unable to create steps')
              .code(500)
          }
          recipe.steps = steps
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

