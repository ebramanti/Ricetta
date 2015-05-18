var Joi = require('joi');

var ONE_DAY_IN_MINUTES = 1440;
var AT_LEAST_2_TAGS = Joi.array().items(Joi.string()).min(2);

var Schema = module.exports = Joi.object().keys({
  author: Joi.string().regex(/[a-zA-Z0-9]{4,30}/).required(),
  image: Joi.string().uri(),
  title: Joi.string().min(4).max(256).required(),
  description: Joi.string().min(0).max(2048),
  // No ingredients for now
  // ingredients: ...,
  cookTime: Joi.number().min(0).max(ONE_DAY_IN_MINUTES),
  prepTime: Joi.number().min(0).max(ONE_DAY_IN_MINUTES),
  // No steps for now
  // steps: ...,
  tags: AT_LEAST_2_TAGS,
  isPrivate: Joi.boolean()
})
