define(function(require, exports, module) {
    var marionette = require('marionette');
    var template = require('hbs!../templates/layouts/home-layout');

    var Recipes = require('app/collections/recipes').Recipes;
    var RecipesView = require('app/views/recipes-view').RecipesView;
    var RecipeViewer = require('app/views/recipe-viewer').RecipeViewer;

    var HomeLayout = marionette.LayoutView.extend({
        template: template,

        regions: {
            list: '#recipe-list',
            viewer: '#recipe-viewer',
            creator: '#recipe-creator'
        },

        ui: {
        },

        events: {
            'click .recipe': 'viewRecipe'
        },

        initialize: function(options) {
            this.session = options.session;
            this.recipes = new Recipes();
        },

        onRender: function() {
            this.showRecipeList();
        },

        showRecipeList: function(options) {
            var personalRecipes = new RecipesView({
                collection: this.recipes
            });
            this.list.show(personalRecipes);
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
    exports.HomeLayout = HomeLayout;
});
