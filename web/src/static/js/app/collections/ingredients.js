define(function(require, exports, module) {

    var backbone = require('backbone');

    var Ingredient = require('app/models/ingredient').Ingredient;

    var Ingredients = Backbone.Collection.extend({
        model: Ingredient,

        initialize: function(options) {
        }
    });

    exports.Ingredients = Ingredient;

});
