var db = require('../connector')
var build = require('../connector').build

module.exports.wipe = function (callback) {
    var query = build([
        'MATCH (n)',
        'OPTIONAL MATCH n-[r]-()',
        'DELETE n, r'
    ])

    db.query(query, function (err, result) {
        if (err) { throw err; }

        callback()
    })
};