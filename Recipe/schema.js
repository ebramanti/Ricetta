var Joi = require('joi')

var StepSchema = require('./step').Schema

var ONE_DAY_IN_MINUTES = 1440

var Schema = module.exports = Joi.object().keys({
  author: Joi.string().regex(/[a-zA-Z0-9]{4,30}/).required(),
  image: Joi.string().uri(),
  title: Joi.string().min(4).max(256).required(),
  description: Joi.string().min(0).max(2048),
  // No ingredients for now
  // ingredients: ...,
  cookTime: Joi.number().min(0).max(ONE_DAY_IN_MINUTES),
  prepTime: Joi.number().min(0).max(ONE_DAY_IN_MINUTES),
  steps: Joi.array().items(
    Joi.object().keys({
        instruction: Joi.string().min(3).max(2048).required(),
        time: Joi.number().min(1).max(ONE_DAY_IN_MINUTES)
    })
  ).min(1).required(),
  tags: Joi.array().items(Joi.string()).min(1),
  isPrivate: Joi.boolean().required()
})
