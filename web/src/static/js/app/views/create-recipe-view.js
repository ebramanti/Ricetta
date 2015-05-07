define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/create-recipe-view');

    var IngredientTemplate = require('hbs!../templates/recipe/create-ingredient-view');
    var StepTemplate = require('hbs!../templates/recipe/create-step-view');
    var TagTemplate = require('hbs!../templates/recipe/create-tag-view');

    var Recipe = require('app/models/recipe').Recipe;

    var CreateRecipeView = marionette.ItemView.extend({
        template: template,
        ui: {
            title: '#title',
            image: '#image',
            preptime: '#preptime',
            preptimeunit: '#preptimeunit',
            cooktime: '#cooktime',
            cooktimeunit: '#cooktimeunit',
            notes: '#notes',
            ingredient: '.ingredient',
            ingredients: '#ingredients-container',
            addingredients: '#add-ingredients',
            steps: '#steps-container',
            addsteps: '#add-steps',
            tags: '#tags-container',
            addtags: '#add-tags',
            private: '#private',
            createButton: '#create'
        },

        events: {
            'click #add-ingredients': 'addIngredients',
            'click #add-steps': 'addSteps',
            'click #add-tags': 'addTags',
            'click #create': 'createRecipe'
        },

        initialize: function(options) {
            this.session = options.session;
            this.collection = options.collection;
        },

        addIngredients: function() {
            this.ui.ingredients.append(IngredientTemplate);
        },

        addSteps: function() {
            this.ui.steps.append(StepTemplate);
        },

        addTags: function() {
            this.ui.tags.append(TagTemplate);
        },

        getIngredients: function() {
            var ingredientResult = [];
            $.each(this.ui.ingredients, function(index, ingredients) {
                $('.ingredient', ingredients).each(function(index, ingredient) {
                    var $ingredient = $(ingredient);
                    var name = $ingredient.find('#name').val();
                    var url = $ingredient.find('#url').val();
                    var amount = $ingredient.find('#amount').val();
                    var amountunit = $ingredient.find('#amountunit').val();
                    ingredientResult.push({
                        name: name,
                        url: url,
                        amount: parseFloat(amount, 10),
                        amountunit: amountunit
                    });
                });
            });
            return ingredientResult;
        },

        getSteps: function() {
            var stepResult = [];
            $.each(this.ui.steps, function(index, steps) {
                $('.step', steps).each(function(index, step) {
                    var $step = $(step);
                    var instruction = $step.find('#instruction').val();
                    var time = $step.find('#time').val();
                    var timeunit = $step.find('#timeunit option:selected').text().slice(0, -3);
                    if (time > 1) {
                        timeunit = timeunit + 's';
                    }
                    stepResult.push({
                        instruction: instruction,
                        time: parseInt(time, 10),
                        timeunit: timeunit
                    });
                });
            });
            return stepResult;
        },

        getTags: function() {
            var tagResult = [];
            $.each(this.ui.tags, function(index, tags) {
                $('.tag', tags).each(function(index, tag) {
                    var name = $(tag).find("input").val();
                    tagResult.push({
                        name: name
                    });
                });
            });
            return tagResult;
        },

        createRecipe: function() {
            // Validate time for cooktime/preptime
            var cooktime = this.ui.cooktime.val();
            var cooktimeunit = this.ui.cooktimeunit.find('option:selected').text();
            cooktimeunit = this.validateTimePlurality(cooktime, cooktimeunit);
            var preptime = this.ui.preptime.val();
            var preptimeunit = this.ui.preptimeunit.find('option:selected').text();
            preptimeunit = this.validateTimePlurality(preptime, preptimeunit);

            var recipe = new Recipe({
                title: this.ui.title.val(),
                image: this.ui.image.val(),
                notes: this.ui.notes.val(),
                ingredients: this.getIngredients(),
                cooktime: parseInt(cooktime, 10),
                cooktimeunit: cooktimeunit,
                preptime: parseInt(preptime, 10),
                preptimeunit: preptimeunit,
                steps: this.getSteps(),
                tags: this.getTags(),
                private: this.ui.private.is(':checked')
            });

            this.collection.create(recipe, {wait: true, error: function(model, res) {
                console.log(res.responseText);
            }});
        },

        validateTimePlurality: function(time, timeunit) {
            var result = timeunit.slice(0, -3);
            if(time === 1) {
                return result;
            } else {
                return result + 's';
            }
        }

    });

    exports.CreateRecipeView = CreateRecipeView;
})
