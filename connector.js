var config = require('config')
var username = config.get('neo4j.username')
var password = config.get('neo4j.password')

var db = require("seraph")({
  server: "http://localhost:7474",
  endpoint: "/db/data",
  user: username,
  pass: password
});

module.exports = db