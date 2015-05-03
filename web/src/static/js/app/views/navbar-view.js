define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/navbar-view')

    var Session = require('app/models/session').Session;

    var NavbarView = marionette.ItemView.extend({
        model: Session,
        template: template,

        ui: {
        },

        events: {

        },

        initialize: function(options) {
            this.model = options.session
        }

    });

    exports.NavbarView = NavbarView;
})
