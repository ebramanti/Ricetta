var util  = require('util'),
    neo4j = require('neo4j-js'),
    frisby = require('frisby');

var helper   = require('./helper');
var accounts = require('./accounts');

describe('Ricetta API test suite', function() {
    helper.test(accounts);
});
