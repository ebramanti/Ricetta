var neo4j = require('neo4j');
var db = new neo4j.GraphDatabase('http://localhost:7474');

// MATCH (n)
// OPTIONAL MATCH (n)-[r]-()
// DELETE n, r

var queries = {
    wipe: [
        "MATCH (n)",
        "OPTIONAL MATCH (n)-[r]-()",
        "DELETE n, r"
    ],
    initialize: "CREATE (pc:PublicCurator {name: \"PublicCurator\"})"
}

var helper = {
    constants: {
        host: 'http://localhost:8228/v1',
        // Routes
        signupRoute: '/signup'
    },
    cleanupAndInit: function() {
        var clean = queries.wipe.join('\n')
        var init  = queries.initialize;

        db.query(clean, function(err, results) {
            // A clean returns 0 rows,
            // throw error if something returns.
            if (results) {
                throw err;
            }
        });

        db.query(init, function(err, results) {
            if (err) {
                console.log("Init: " + err);
                throw err;
            }
        });
    },
    cleanup: function() {
        var query = queries.wipe.join('\n');

        db.query(query, function(err, results) {
            if (err) {
                throw err;
            }
        });
    }
};

module.exports = helper;
