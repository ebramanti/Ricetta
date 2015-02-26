var neo4j = require('neo4j-js');

var helper = {
    constants: {
        host: 'http://localhost:8228/v1'
    },
    test: function(spec, cleanup) {
        for (var key in spec) {
            if (spec.hasOwnProperty(key)) {
                var module = spec[key];
                describe(module.description, function() {
                    afterEach(cleanup);
                    for (var test in module) {
                        if (module.hasOwnProperty(test) && test !== 'description') {
                            var currentTest = module[test]
                            it(currentTest.description, function() {
                                currentTest.request.toss();
                            });
                        }
                    }
                });
            }
        }
    },
    cleanup: function() {
        //All testing happens locally, therefore harcoded Neo URL.
        neo4j.connect('http://localhost:7474/db/data/', function (err, graph) {
            if (err) {
                throw err;
            } else {
                var query = [
                    "MATCH (n)",
                    "OPTIONAL MATCH (n)-[r]-()",
                    "DELETE n, r"
                ];

                graph.query(query.join('\n'), function(err, results) {
                    if (err) {
                        console.log(err);
                        console.log(err.stack);
                    } else {
                        console.log("Database cleaned, teardown complete.")
                    }
                });
            }
        });
    }
};

module.exports = helper;
