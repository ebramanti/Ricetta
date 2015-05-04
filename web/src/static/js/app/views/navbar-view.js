define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/navbar-view')

    var Session = require('app/models/session').Session;

    var NavbarView = marionette.ItemView.extend({
        model: Session,
        template: template,

        ui: {
            logout: '#logout'
        },

        events: {
            'click #logout': 'onLogout'
        },

        initialize: function(options) {
            this.model = options.session;
            this.session = options.session;
        },

        onLogout: function(event) {
            event.preventDefault();
            this.session.logout({
                reload: true
            });
            window.location.replace('/');
        }

    });

    exports.NavbarView = NavbarView;
})
