var frisby = require('frisby'),
    constants  = require('../helper').constants,
    cleanupAndInit = require('../helper').cleanupAndInit;
    cleanup = require('../helper').cleanup;

var frisbyCleanup = function() {
    cleanupAndInit();
}

var signup = function() {
    describe('Signup Tests', function() {
        afterEach(cleanupAndInit);
        afterAll(cleanup);

        frisby.create('tests a missing handle').post(constants.host + constants.signupRoute, {
            handle: "",
            email: "socash@ricetta.io",
            password: "12345678",
            confirmpassword: "12345678"
        }, {json: true}).expectJSON({
            reason: ["field handle is invalid: Required field for signup"]
        }).expectStatus(400).toss();

        frisby.create('tests a missing email').post(constants.host + constants.signupRoute, {
            handle: "socash",
            email: "",
            password: "12345678",
            confirmpassword: "12345678"
        }, {json: true}).expectJSON({
            reason: ["field email is invalid: Required field for signup"]
        }).expectStatus(400).toss();

        frisby.create('tests a missing password').post(constants.host + constants.signupRoute, {
            handle: "socash",
            email: "socash@ricetta.io",
            password: "",
            confirmpassword: "12345678"
        }, {json: true}).expectJSON({
            reason: ["field password is invalid: Required field for signup"]
        }).expectStatus(400).toss();

        frisby.create('tests a missing confirmpassword').post(constants.host + constants.signupRoute, {
            handle: "socash",
            email: "socash@ricetta.io",
            password: "12345678",
            confirmpassword: ""
        }, {json: true}).expectJSON({
            reason: ["field confirmpassword is invalid: Required field for signup"]
        }).expectStatus(400).toss();

        frisby.create('tests password too short').post(constants.host + constants.signupRoute, {
            handle: "socash",
            email: "socash@ricetta.io",
            password: "1234567",
            confirmpassword: "1234567"
        }, {json: true}).expectJSON({
            reason: ["field password is invalid: Too short, minimum length is 8"]
        }).expectStatus(400).toss();

        frisby.create('tests password too long').post(constants.host + constants.signupRoute, {
            handle: "socash",
            email: "socash@ricetta.io",
            password: "123456789012345678901234567890123456789012345678901",
            confirmpassword: "123456789012345678901234567890123456789012345678901"
        }, {json: true}).expectJSON({
            reason: ["field password is invalid: Too long, maximum length is 50"]
        }).expectStatus(400).toss();

        frisby.create('tests mismatch of password and confirmation').post(constants.host + constants.signupRoute, {
            handle: "socash",
            email: "socash@ricetta.io",
            password: "12345678",
            confirmpassword: "123456789"
        }, {json: true}).expectJSON({
            reason: "Passwords do not match"
        }).expectStatus(403).toss();

        frisby.create('fails if handle already exists').post(constants.host + constants.signupRoute, {
            handle: "socash",
            email: "socash@ricetta.io",
            password: "12345678",
            confirmpassword: "12345678"
        }, {json: true}).expectJSON({
            email: "socash@ricetta.io",
            handle: "socash",
            response: "Signed up a new user!"
        }).after(function() {
            frisby.create('fails on second user').post(constants.host + constants.signupRoute, {
                handle: "socash",
                email: "socash2@ricetta.io",
                password: "123456789",
                confirmpassword: "123456789"
            }, {json: true}).expectJSON({
                reason: "Sorry, handle or email is already taken"
            }).expectStatus(409).toss();
        }).expectStatus(201).toss();

        // frisby.create('fails if email already exists').post(constants.host + constants.signupRoute, {
        //     handle: "socash",
        //     email: "socash@ricetta.io",
        //     password: "12345678",
        //     confirmpassword: "12345678"
        // }, {json: true}).expectJSON({
        //     email: "socash@ricetta.io",
        //     handle: "socash",
        //     response: "Signed up a new user!"
        // }).after(function() {
        //     frisby.create('fails on second user').post(constants.host + constants.signupRoute, {
        //         handle: "socash2",
        //         email: "socash@ricetta.io",
        //         password: "123456789",
        //         confirmpassword: "123456789"
        //     }, {json: true}).expectJSON({
        //         reason: "Sorry, handle or email is already taken"
        //     }).expectStatus(409).toss();
        // }).expectStatus(201).toss();

        // frisby.create('tests a good signup').post(constants.host + constants.signupRoute, {
        //     handle: "socash",
        //     email: "socash@ricetta.io",
        //     password: "12345678",
        //     confirmpassword: "12345678"
        // }, {json: true}).expectJSON({
        //     email: "socash@ricetta.io",
        //     handle: "socash",
        //     response: "Signed up a new user!"
        // }).inspectJSON().expectStatus(201).toss();

    });
}

module.exports = signup;
