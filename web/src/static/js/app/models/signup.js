define(function(require, exports, module) {

    var backbone = require('backbone');

    var Signup = Backbone.Model.extend({
        /*validate: function(attrs, options) {
            if (attrs.password !== attrs.confirmpassword) {
                return "Must have same password/password confirmation";
            }
        },*/
        url: '/v1/signup',
        defaults: {
            handle: null,
            email: null,
            password: null,
            confirmpassword: null
        },

        authenticate: function() {
            this.save({}, {
                success: function(model, response) {
                    return true;
                },
                error: function(model, response) {
                    console.log(response.responseText)
                }
            })
        },

        initialize: function() {

        }
    });

    exports.Signup = Signup;

});
