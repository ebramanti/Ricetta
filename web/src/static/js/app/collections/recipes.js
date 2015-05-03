define(function(require, exports, module) {

    var backbone = require('backbone');

    var Recipe = require('app/models/recipe').Recipe;

    var Recipes = Backbone.Collection.extend({
        url: 'v1/recipes',
        model: Recipe,

        // parse: function(response) {
        //     return response;
        // },

        initialize: function() {

        }
    });

    exports.Recipes = Recipes;

});
