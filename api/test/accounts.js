var frisby = require('frisby'),
    constants  = require('./helper').constants;

// Set up Frisby Tests

var accounts = {};

// Tests for Signup
var signup = {};
signup.description = 'Signup Tests';

signup.goodUser = {
    description: 'tests a good signup',
    request: frisby.create(this.description).post(constants.host + '/signup', {
        handle: "socash",
        email: "socash@ricetta.io",
        password: "12345678",
        confirmpassword: "12345678"
    }, {json: true}).expectJSON({
        email: "socash@ricetta.io",
        handle: "socash",
        response: "Signed up a new user!"
    }).expectStatus(201)
};

// signup.badHandle = {
//     description: 'tests a missing handle',
//     request: frisby.create(this.description).post(constants.host + '/signup', {
//         handle: "",
//         email: "socash@ricetta.io",
//         password: "12345678",
//         confirmpassword: "12345678"
//     }, {json: true}).expectJSON({
//         reason: "field handle is invalid: Required field for signup"
//     }).expectStatus(400)
// };

// Add specific test modules to accounts object
accounts.signup = signup;

module.exports = accounts;

