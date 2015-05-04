define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/curator-view')
    var RecipeViewer = require('app/views/recipe-viewer').RecipeViewer;
    var Recipe = require('app/models/recipe').Recipe;

    var RecipeModal = marionette.CompositeView.extend({
        childView: RecipeViewer,
        childViewContainer: '#curator',
        template: template,

        ui: {
        },

        initialize: function(options) {
            this.collection = options.collection;
            this.session = options.session;
            this.collection.fetch({
                success: function(res) {
                    console.log(res);
                }
            });
        }

    });

    exports.RecipeModal = RecipeModal;
})
