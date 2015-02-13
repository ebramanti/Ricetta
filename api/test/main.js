var util  = require('util'),
    spawn = require('child_process').spawn,
    api   = spawn("go", ["run", "main.go"], {cwd: "../"});

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
