var db = require('../connector')
var Code = require('code')
var expect = Code.expect
var Lab = require('lab')
var lab = exports.lab = Lab.script()

lab.experiment('The database', function (done) {

  lab.test('is actually on and connected to', function (done) {
    db.save({ type: 'TestNode' }, 'TEST', function (err, node) {
      expect(err ? err.code : null).to.be.null()
      done()
    })
  })

})
