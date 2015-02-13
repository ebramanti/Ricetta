var util  = require('util'),
    neo4j = require('neo4j-js'),
    spawn = require('child_process').spawn,
    api; // This is the child process that starts the Go API.

function SetUpSuite() {
    api = spawn("go", ["run", "main.go", "-testing=true"], {cwd: "../"});
    api.stdout.on('data', function (data) {
        console.log("API: " + data);
    });

    api.stderr.on('data', function (data) {
        console.log("API: " + data);
    });

    api.on('exit', function (code) {
        console.log('child process exited with code ' + code);
    });
}

function TearDownSuite() {
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
    api.kill();
}

(function test() {
    SetUpSuite();
})();
