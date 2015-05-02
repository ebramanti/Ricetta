define(function(require, exports, module) {

    var backbone = require('backbone');

    var Recipe = Backbone.Model.extend({
        url: 'v1/recipes',
        defaults: {
            id: null,
            created: null,
            last_modified: null,
            author: null,
            url: null,
            title: null,
            notes: null,
            ingredients: null,
            cooktime: null,
            cooktimeunit: null,
            preptime: null,
            preptimeunit: null,
            steps: null,
            tags: null,
            private: null
        },

        initialize: function() {

        }
    });

    exports.Recipe = Recipe;

})
