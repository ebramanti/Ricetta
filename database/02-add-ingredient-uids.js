var db = require('../connector')
var shortid = require('shortid')


db.nodesWithLabel('Ingredient', function (err, results) {
  if (err) { throw err; }
  results.forEach(function (ingredient) {
    ingredient.uid = shortid.generate() // e.g. `23TplPdS`
    db.save(ingredient, function (err, node) {
      if (err) { throw err; }
    })
  })
  console.log('02: Done adding uid\'s to ingredients')
})