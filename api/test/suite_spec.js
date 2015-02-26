var util  = require('util'),
    frisby = require('frisby');

var helper   = require('./helper');
var cleanup  = helper.cleanup;
var accounts = require('./accounts');

describe('Ricetta API test suite', function() {
    describe('Ricetta Test Modules', function() {
        helper.test(accounts, cleanup);
    });
    // it('tears down the tests', function() {
    //     cleanup();
    // });
});
