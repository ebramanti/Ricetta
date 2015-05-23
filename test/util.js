var db = require('../connector')
var build = require('../connector').build


module.exports.setUpLab = function (lab) {
    lab.before(function (done) {
      wipe(done)
    })
    lab.afterEach(function (done) {
      wipe(done)
    })
}


function wipe (callback) {
    var query = build([
        'MATCH (n)',
        'OPTIONAL MATCH n-[r]-()',
        'DELETE n, r'
    ])

    db.query(query, function (err, result) {
        if (err) { throw err; }

        callback(err, null)
    })
};
