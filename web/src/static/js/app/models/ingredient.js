define(function(require, exports, module) {

    var backbone = require('backbone');

    var Ingredient = Backbone.Model.extend({
        defaults: {
            name: null,
            amount: null,
            amountunit: null,
            url: null
        },

        initialize: function() {

        }
    });

    exports.Ingredient = Ingredient;

});
