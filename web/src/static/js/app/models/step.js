define(function(require, exports, module) {

    var backbone = require('backbone');

    var Step = Backbone.Model.extend({
        defaults: {
            instruction: null,
            time: null,
            timeunit: null
        },

        initialize: function() {

        }
    });

    exports.Step = Step;

});
