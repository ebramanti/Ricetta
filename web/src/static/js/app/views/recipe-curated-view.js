define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/recipe-curated-view');
    var Recipe = require('app/models/recipe').Recipe;

    var RecipeCuratedView = marionette.ItemView.extend({
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

    exports.RecipeCuratedView = RecipeCuratedView;
})
