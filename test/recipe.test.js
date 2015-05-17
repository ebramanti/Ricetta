var assert = require('assert')
var Lab = require('lab')
var lab = exports.lab = Lab.script()

var app = require('../app')

lab.experiment('The /recipe endpoint', function (done) {


  lab.test('No data to /recipe fails', function (done) {
    var options = {
      method: 'POST',
      url: '/recipes'
    }
    app.inject(options, function (res) {
      assert.equal(res.statusCode, 400)
      done()
    })
  })


  lab.test('Wrong /recipe route', function (done) {
    var options = {
      method: 'POST',
      url: '/reciPIES'
    }
    app.inject(options, function (res) {
      assert.equal(res.statusCode, 404)
      done()
    })
  })


  lab.test('Good /recipe post', function (done) {
    var options = {
      method: 'POST',
      url: '/recipes',
      payload: {
        "title": "Old Fashioned Pancakes",
        "author": "SoCash",
        "notes": "Tastes almost as good as Grandma's recipe.",
        "cookTime": 15,
        "prepTime": 5,
        "tags": [
          "trill",
          "maple"
        ],
        "isPrivate": false
      }
    }
    app.inject(options, function (res) {
      assert.equal(res.statusCode, 201)
      done()
    })
  })


})