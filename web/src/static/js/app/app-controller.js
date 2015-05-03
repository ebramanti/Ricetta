define(function(require, exports, module) {
    var backbone = require('backbone');
    var marionette = require('marionette');
    var app = require('app/app');
    var $ = require('jquery');

    var Session = require('app/models/session').Session;

    var LandingLayout = require('app/layouts/landing-layout').LandingLayout;
    var HomeLayout = require('app/layouts/home-layout').HomeLayout;

    var SignupView = require('app/views/signup-view').SignupView;
    var FooterView = require('app/views/footer-view').FooterView;
    var LoginView = require('app/views/login-view').LoginView;
    var NavbarView = require('app/views/navbar-view').NavbarView


    var AppController = marionette.Controller.extend({
        initialize: function() {
            this.app = app;
            this.app.session = new Session({}, {
                remote: false
            });
            // Logic for auth check.
            if (this.app.session.has('token')) {
                console.log("User logged in");
                $.ajaxSetup({
                    headers: {'Authorization' : this.app.session.get('token')}
                })
                this.app.session = app.session;
                // user is authed, redirect home
                this.showHomeLayout();
            };
        },
        // Needed for AppRouter to initialize index route.
        index: function() {
            console.log(this.app.session)
            if (this.app.session.hasAuth()) {
                this.showHomeLayout();
            } else {
                this.app.mainRegion.show(new LandingLayout({
                    session: this.app.session
                }));
            }
        },

        showHomeLayout:function() {
            if (!this.app.session.hasAuth()) {
                this.index();
            } else {
                this.app.headerRegion.show(new NavbarView({
                    session: this.app.session
                }));
                this.app.mainRegion.show(new HomeLayout({
                    session: this.app.session
                }));
            }
        }
    });
    exports.AppController = AppController;
});
