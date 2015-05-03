define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/recipe-view');
    var Recipe = require('app/models/recipe').Recipe;

    var RecipeView = marionette.ItemView.extend({
        model: Recipe,
        template: template,

        ui: {

        },

        events: {
        },

        initialize: function(options) {
            this.model = options.model
            this.session = options.session;
        }

    });

    exports.RecipeView = RecipeView;
})
