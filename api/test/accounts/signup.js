var frisby = require('frisby'),
    constants  = require('../helper').constants,
    cleanup = require('../helper').cleanup;

var signup = function() {
    describe('Signup Tests', function() {
        afterEach(cleanup)
        frisby.create('tests a good signup').post(constants.host + '/signup', {
            handle: "socash",
            email: "socash@ricetta.io",
            password: "12345678",
            confirmpassword: "12345678"
        }, {json: true}).expectJSON({
            email: "socash@ricetta.io",
            handle: "socash",
            response: "Signed up a new user!"
        }).expectStatus(201).toss();

        // frisby.create('tests a missing handle').post(constants.host + '/signup', {
        //     handle: "",
        //     email: "socash@ricetta.io",
        //     password: "12345678",
        //     confirmpassword: "12345678"
        // }, {json: true}).expectJSON({
        //     reason: "field handle is invalid: Required field for signup"
        // }).expectStatus(400).toss();
    });
}

module.exports = signup;
