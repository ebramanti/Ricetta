var Code = require('code')
var expect = Code.expect
var Lab = require('lab')
var lab = exports.lab = Lab.script()

var app = require('../app')

lab.experiment('The /recipe endpoint', function (done) {
  lab.experiment('When posted to', function (done) {

    lab.test('no data to /recipe fails', function (done) {
      var options = {
        method: 'POST',
        url: '/recipes'
      }

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
      var options = {
        method: 'POST',
        url: '/recipes',
        payload: {
          "title": "Old Fashioned Pancakes",
          "author": "SoCash",
          "notes": "Tastes almost as good as Grandma's recipe.",
          "cookTime": 15,
          "prepTime": 5,
          "tags": [ "trill", "maple" ],
          "isPrivate": false
        }
      }

      app.inject(options, function (res) {
        var body = JSON.parse(res.payload)

        expect(res.statusCode).to.not.equal(500)
        expect(res.statusCode).to.equal(201)
        expect(body).to.not.include('error')
        expect(body).to.be.an.object()

        done()
      })
    })

  })

  lab.experiment('GET route', function (done) {

    lab.test('returns a list of public recipes', function (done) {
      var options = {
        method: 'GET',
        url: '/recipes'
      }

      app.inject(options, function (res) {
        var body = JSON.parse(res.payload)

        expect(res.statusCode).to.equal(200)
        expect(body).to.be.an.array()

        done()
      })
    })

  })
})