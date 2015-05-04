define(function(require, exports, module) {

    var backbone = require('backbone');

    var Recipe = require('app/models/recipe').Recipe;

    var Recipes = Backbone.Collection.extend({
        url: 'v1/recipes',
        model: Recipe,

        initialize: function(options) {
            var object = options || {curator: false}
            if (object.curator) {
                this.url = this.url + '?curator'
            }
        }
    });

    exports.Recipes = Recipes;

});
