define(function(require, exports, module) {
    var marionette = require('marionette');
    var template = require('hbs!../templates/signup-view');

    var Signup = require('app/models/signup').Signup;
    var Login = require('app/models/login').Login;

    var SignupView = marionette.ItemView.extend({
        template: template,

        //takes the div marionette creates and give it a class named mainContainer.
        tagName: "div",
        className: "mainContainer",
        ui: {
            handle: '#handle',
            email: '#input-email',
            pass: '#pass1',
            confirmPass: '#pass2'
        },

        events: {
            'click #signup': 'onFormConfirm'
        },

        initialize: function(options) {
            this.session = options.session;
        },

        onFormConfirm: function(event) {
            event.preventDefault();
            var req = new Signup({
                handle: this.ui.handle.val(),
                email: this.ui.email.val(),
                password: this.ui.pass.val(),
                confirmpassword: this.ui.confirmPass.val()
            });
            req.authenticate();
            console.log("Success on signup");
            debugger;
            console.log("Logging in");
            var login = new Login({
                handle: this.ui.handle.val(),
                password: this.ui.pass.val(),
                session: this.session
            })
            login.authenticate();
            login.clear();
        }

    });

    exports.SignupView = SignupView;
})
