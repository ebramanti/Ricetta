define(function(require, exports, module) {
    var marionette = require('marionette');
    var template = require('hbs!../templates/layouts/curator-layout');

    var Recipes = require('app/collections/recipes').Recipes;
    var RecipesView = require('app/views/recipes-view').RecipesView;
    var RecipeViewer = require('app/views/recipe-viewer').RecipeViewer;

    var CuratorView = require('app/views/curator-view').CuratorView;

    var CuratorLayout = marionette.LayoutView.extend({
        template: template,

        regions: {
            curator: '#curator',
            viewer: '#curatorViewer'
        },

        ui: {
            curatorViewer: '#curatorViewer'
        },

        events: {
            'click .recipe': 'viewRecipe'
        },

        initialize: function(options) {
            this.session = options.session;
            this.curated = new Recipes({
                curator: true
            });
        },

        onRender: function() {
            this.ui.curatorViewer.hide();
            this.showCuratedList();
        },

        showCuratedList: function(options) {
            var curatedRecipes = new CuratorView({
                collection: this.curated
            });
            this.curator.show(curatedRecipes);
        },

        viewRecipe: function(event) {
            if (this.curatorViewerHidden()) {
                this.ui.curatorViewer.show();
            }
            var id = $(event.currentTarget).attr('id');
            var currentRecipe = this.curated.get(id);
            var currentRecipeView = new RecipeViewer({
                model: currentRecipe
            });
            this.viewer.show(currentRecipeView);
        },
        curatorViewerHidden: function() {
            return this.ui.curatorViewer.is(':hidden');
        }

    });
    exports.CuratorLayout = CuratorLayout;
});
