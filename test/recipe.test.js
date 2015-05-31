var Code = require('code')
var expect = Code.expect
var Lab = require('lab')
var lab = exports.lab = Lab.script()
var async = require('async')

var app = require('../app')

var wipe = require('./util').wipe

lab.afterEach(function (done) {
  wipe(done)
})

var GET_RECIPES = {
  method: 'GET',
  url: '/recipes'
}
var POST_RECIPE = function(payload) {
  return {
    method: 'POST',
    url: '/recipes',
    payload: payload
  }
}

lab.experiment('The /recipe endpoint', function (done) {
  lab.experiment('When posted to', function (done) {

    lab.test('no data to /recipe fails', function (done) {
      var options = POST_RECIPE({})

      app.inject(options, function (res) {
        expect(res.statusCode).to.equal(400)
        expect(res.payload).to.include('error')
        expect(res.payload).to.include('message')
        expect(res.payload).to.include('validation')
        done()
      })
    })

    lab.test('wrong /recipe route fails', function (done) {
      var options = {
        method: 'POST',
        url: '/reciPIES'
      }

      app.inject(options, function (res) {
        expect(res.statusCode).to.equal(404)
        done()
      })
    })

    lab.test('can create a recipe', function (done) {
      var options = POST_RECIPE({
        "title": "Old Fashioned Pancakes",
        "author": "SoCash",
        "description": "Tastes almost as good as Grandma's recipe.",
        "cookTime": 15,
        "prepTime": 5,
        "steps": [
          {
            "instruction": "In a large bowl, sift together the flour, baking powder, salt and sugar. Make a well in the center and pour in the milk, egg and melted butter; mix until smooth.",
            "time": 10
          },
          {
            "instruction": "Heat a lightly oiled griddle or frying pan over medium high heat. Pour or scoop the batter onto the griddle, using approximately 1/4 cup for each pancake. Brown on both sides and serve hot."
          }
        ],
        "tags": [ "trill", "maple" ],
        "isPrivate": false
      })

      app.inject(options, function (res) {
        var body = JSON.parse(res.payload)

        expect(res.statusCode).to.not.equal(500)
        expect(res.statusCode).to.equal(201)
        expect(body).to.not.include('error')
        expect(body).to.be.an.object()
        expect(body.steps.length).to.equal(2)

        done()
      })
    })

  })

  lab.experiment('GET route', function (done) {

    lab.test('returns a list of public recipes', function (done) {
      var options = GET_RECIPES

      app.inject(options, function (res) {
        var body = JSON.parse(res.payload)

        expect(res.statusCode).to.equal(200)
        expect(body).to.be.an.array()
        expect(body.length).to.equal(0)

        done()
      })
    })

    lab.test('can return a couple recipes', function (done) {
      var options = POST_RECIPE({
        "title": "Old Fashioned Pancakes",
        "author": "AsyncMan",
        "description": "Tastes almost as good as Grandma's recipe.",
        "cookTime": 15,
        "prepTime": 5,
        "steps": [
          {
            "instruction": "In a large bowl, sift together the flour, baking powder, salt and sugar. Make a well in the center and pour in the milk, egg and melted butter; mix until smooth.",
            "time": 10
          },
          {
            "instruction": "Heat a lightly oiled griddle or frying pan over medium high heat. Pour or scoop the batter onto the griddle, using approximately 1/4 cup for each pancake. Brown on both sides and serve hot.",
            "time": 10
          }
        ],
        "tags": [ "trill", "maple" ],
        "isPrivate": false
      })

      async.series([
        function addRecipe (callback) {
          app.inject(options, function (res) {
            expect(res.statusCode).to.equal(201)
            callback()
          })
        },
        function checkRecipe (callback) {
          app.inject(GET_RECIPES, function (res) {
            var body = JSON.parse(res.payload)

            expect(res.statusCode).to.equal(200)
            expect(body.length).to.equal(1)
            callback()
          })
        }
      ], function (error, results) {
        done()
      })
    })

  })
})
