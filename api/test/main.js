var util  = require('util'),
    spawn = require('child_process').spawn,
    api; // This is the child process that starts the Go API.

function SetUpSuite() {
    api = spawn("go", ["run", "main.go"], {cwd: "../"});
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
        }
        // do something with the graph
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

    });
    api.kill();
}

(function test() {
    SetUpSuite();
})();
