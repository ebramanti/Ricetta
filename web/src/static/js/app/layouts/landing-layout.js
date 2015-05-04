define(function(require, exports, module) {
    var marionette = require('marionette');
    var template = require('hbs!../templates/layouts/landing-layout')
    var SignupView = require('app/views/signup-view').SignupView;
    var LoginView = require('app/views/login-view').LoginView;

    var Recipes = require('app/collections/recipes').Recipes;

    var CuratorView = require('app/views/curator-view').CuratorView;
    var Login = require('app/models/login').Login;
    var LandingLayout = marionette.LayoutView.extend({
        template: template,

        regions: {
            signupLoginArea: '#signupLoginArea',
            curator: '#curatorArea'
        },

        ui: {
            jumbotron: '.jumbotron',
            existingUser: '#goToLogin',
            newUser: '#goToSignup'
        },

        events: {
            'click #login': 'showLoginForm',
            'click #signup': 'showSignupForm',
            'click .recipe': 'viewRecipe'
        },
        initialize: function(options) {
            this.session = options.session;
            this.recipes = new Recipes({
                curator: true
            })
        },

        onRender: function(options) {
            var curatorView = new CuratorView({
                collection: this.recipes
            });
            this.curator.show(curatorView);
        },

        showLoginForm: function(options) {
            var loginView = new LoginView({
                session: this.session
            });
            this.ui.jumbotron.hide();
            this.signupLoginArea.show(loginView);

        },

        showSignupForm: function(options) {
            var signupView = new SignupView({
                session: this.session
            });
            this.ui.jumbotron.hide();
            this.signupLoginArea.show(signupView);
        },

        viewRecipe: function(event) {
            var id = $(event.currentTarget).attr('id');
            var currentRecipe = this.recipes.get(id);
            var currentRecipeView = new RecipeViewer({
                model: currentRecipe
            });
            this.viewer.show(currentRecipeView);
        }

    });
    exports.LandingLayout = LandingLayout;
})
