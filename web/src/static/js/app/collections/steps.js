define(function(require, exports, module) {

    var backbone = require('backbone');

    var Step = require('app/models/Steps').Step;

    var Steps = Backbone.Collection.extend({
        model: Step,

        initialize: function(options) {
        }
    });

    exports.Steps = Steps;

});
