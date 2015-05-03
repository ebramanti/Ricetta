define(function(require, exports, module) {
    var marionette = require('marionette');
    var template = require('hbs!../templates/layouts/home-layout');

    var Recipes = require('app/collections/recipes').Recipes;
    var RecipesView = require('app/views/recipes-view').RecipesView;

    var HomeLayout = marionette.LayoutView.extend({
        template: template,

        regions: {
            list: '#recipe-list',
            viewer: '#recipe-viewer',
            creator: '#recipe-creator'
        },

        initialize: function(options) {
            this.session = options.session;
        },

        onRender: function() {
            this.showRecipeList();
        },

        showRecipeList: function(options) {
            var recipes = new Recipes();
            var personalRecipes = new RecipesView({
                collection: recipes
            });
            this.list.show(personalRecipes);
        }
    });
    exports.HomeLayout = HomeLayout;
});
