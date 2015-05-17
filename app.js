var Hapi = require('hapi')
var port = '8000'
var server = new Hapi.Server()
server.connection({
	host: 'localhost',
	port: port
})
server.route(require('./routes'))

var connector = require('./connector')

server.start()

console.log("Started Ricetta API on port " + port)
