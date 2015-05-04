define(function(require, exports, module) {
    var marionette = require('marionette');
    var template = require('hbs!../templates/layouts/home-layout');

    var Recipes = require('app/collections/recipes').Recipes;
    var RecipesView = require('app/views/recipes-view').RecipesView;
    var RecipeViewer = require('app/views/recipe-viewer').RecipeViewer;
    var RecipeViewerDefault = require('app/views/recipe-viewer-default').RecipeViewerDefault;

    var CreateRecipeView = require('app/views/create-recipe-view').CreateRecipeView;

    var HomeLayout = marionette.LayoutView.extend({
        template: template,

        regions: {
            list: '#recipe-list',
            viewer: '#recipe-viewer'
        },

        ui: {
        },

        events: {
            'click .recipe': 'viewRecipe',
            'click .create-recipe': 'createRecipe',
            'click #cancel': 'resetViewer'
        },

        initialize: function(options) {
            this.session = options.session;
            this.recipes = new Recipes();
            this.curator = new Recipes({
                curator: true
            });
        },

        onRender: function() {
            this.showRecipeList();
        },

        showRecipeList: function(options) {
            var personalRecipes = new RecipesView({
                collection: this.recipes
            });
            this.list.show(personalRecipes);
            this.viewer.show(new RecipeViewerDefault());
        },
        viewRecipe: function(event) {
            var id = $(event.currentTarget).attr('id');
            var currentRecipe = this.recipes.get(id);
            var currentRecipeView = new RecipeViewer({
                model: currentRecipe
            });
            this.viewer.show(currentRecipeView);
        },
        createRecipe: function() {
            console.log("Success")
            this.viewer.show(new CreateRecipeView({
                collection: this.recipes
            }));
        },
        resetViewer: function() {
            this.viewer.show(new RecipeViewerDefault());
        }
    });
    exports.HomeLayout = HomeLayout;
});
