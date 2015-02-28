var neo4j = require('neo4j');
var db = new neo4j.GraphDatabase('http://localhost:7474');

var helper = {
    constants: {
        host: 'http://localhost:8228/v1',
        // Routes
        signupRoute: '/signup'
    },
    cleanup: function() {
        var query = [
            "MATCH (n)",
            "OPTIONAL MATCH (n)-[r]-()",
            "DELETE n, r"
        ];

        db.query(query.join('\n'), function(err, results) {
            if (err) {
                throw err;
            }
        });
    }
};

module.exports = helper;
