define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/navbar-view')

    var NavbarView = marionette.ItemView.extend({
        template: template,

        ui: {
        },

        events: {

        },

        initialize: function(options) {
        }

    });

    exports.NavbarView = NavbarView;
})
