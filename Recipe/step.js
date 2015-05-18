var Joi = require('joi')
var async = require('async')
var db = require('../connector')

var Step = module.exports = {
  CreateSteps: function(recipe, steps, next) {
    var stepQueries = []
    // First step
    stepQueries.push(function (cb) {
      async.waterfall([
        function createStepNode (callback) {
          db.save(steps[0], function (err, node) {
            if (err) { return callback(err, null); }
            callback(null, node)
          })
        },
        function linkStepNode (node, callback) {
          db.relate(recipe, 'FIRST_STEP', node, {}, function (err, relationship) {
            if (err) { return callback(err, null); }
            callback(null, node)
          })
        }
      ], function (error, node) {
        cb(error, node)
      })
    })
    // More steps
    for (var i = 1; i < steps.length; i++) {
      stepQueries.push(function (previousNode, cb) {
        async.waterfall([
          function createStepNode (callback) {
            db.save(steps[i], function (err, node) {
              if (err) { return callback(err, null); }
              callback(null, node)
            })
          },
          function linkStepNode (node, callback) {
            db.relate(previousNode, 'NEXT', node, {}, function (err, relationship) {
              if (err) { return callback(err, null); }
              callback(null, node)
            })
          }
        ], function (error, node) {
          cb(error, node)
        })
      })
    }

    async.waterfall(stepQueries, function (error, results) {
      if (error) {
        return reply('Unable to create steps').code(500)
      }
      next()
    })
  }
}
