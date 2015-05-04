define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/recipe-viewer-default');

    var RecipeViewerDefault = marionette.ItemView.extend({
        template: template
    });

    exports.RecipeViewerDefault = RecipeViewerDefault;
})
