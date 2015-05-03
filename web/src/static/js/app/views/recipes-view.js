define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/recipe-view')
    var RecipeView = require('app/views/recipe-view').RecipeView;
    var Recipe = require('app/models/recipe').Recipe;

    var RecipesView = marionette.CollectionView.extend({
        childView: RecipeView,
        template: template,

        ui: {
        },

        initialize: function(options) {
            this.collection = options.collection;
            this.session = options.session;
            this.collection.fetch({
                success: function(res) {
                    console.log(res);
                },
                error: function(collection, res) {
                    console.log(res.responseText);
                }
                // data: $.param({ all: true})
            });
        }

    });

    exports.RecipesView = RecipesView;
})
