var Hapi = require('hapi')
var port = '8000'

var connector = require('./connector')
var server = new Hapi.Server()

server.connection({
	host: 'localhost',
	port: port
})

server.route(require('./routes'))

server.start(function() {
    console.log("Started Ricetta API on port " + port)
})

module.exports = server
