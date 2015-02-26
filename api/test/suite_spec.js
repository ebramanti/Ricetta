var util  = require('util'),
    neo4j = require('neo4j-js'),
    spawn = require('child_process').spawn,
    frisby = require('frisby');

describe('Ricetta API test suite', function() {
    var api; // This is the child process that starts the Go API.

    it('sets up the test suite', function() {
        api = spawn("go", ["run", "main.go", "-testing=true"], {cwd: "../"});
        api.stdout.on('data', function (data) {
            console.log("API: " + data);
        });

        api.stderr.on('data', function (data) {
            console.log("API: " + data);
            api.kill();
        });

        api.on('exit', function (code) {
            console.log('child process exited with code ' + code);
        });
    });

    describe('run the API tests', function() {
        afterEach(function() {
            // All testing happens locally, therefore harcoded Neo URL.
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
        });

        // // tests here
        // frisby.create('Empty Handle Signup')
        // .post('http://localhost:8228/v1' + '/signup', {
        //  handle: '',
        //  email: 'socash@ricetta.io',
        //  password: '12345678',
        //  confirmpassword: '12345678'
        // }, {json: true})
        // .expectJSON({
        // args: {
        //     reason: 'field handle is invalid: Required field for signup'
        // }
        // })
        // .expectStatus(400)
        // .toss();
    });

    it('tears down the test suite', function() {
        api.kill();
    });

});
