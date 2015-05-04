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
            curator: '#curator'
        },

        ui: {
        },

        initialize: function(options) {
            this.session = options.session;
            this.curated = new Recipes({
                curator: true
            });
        },

        onRender: function() {
            this.showCuratedList();
        },

        showCuratedList: function(options) {
            var curatedRecipes = new CuratorView({
                collection: this.curated
            });
            this.curator.show(curatedRecipes);
        },
    });
    exports.CuratorLayout = CuratorLayout;
});
