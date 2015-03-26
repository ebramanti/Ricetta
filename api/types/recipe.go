package types

import (
	"fmt"
)

//
// Recipe Creation Types
//

type Recipe struct {
	Title        string       `json:"title" validate:"recipetitle"`
	Notes        string       `json:"notes" validate:"recipenotes"`
	Ingredients  []Ingredient `json:"ingredients"`
	CookTime     int          `json:"cooktime" validate:"time"`
	CookTimeUnit string       `json:"cooktimeunit" validate:"timeunit"`
	PrepTime     int          `json:"preptime" validate:"time"`
	PrepTimeUnit string       `json:"prepunit" validate:"timeunit"`
	Steps        []Step       `json:"steps"`
	Tags         []Tag        `json:"tags"`
	Public       bool         `json:"public"`
}

//
// Recipe Component Types
//

type Ingredient struct {
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
	AmountUnit string `json:"amountunit"`
	URL        string `json:"url"`
}

type Step struct {
	Instruction string `json:"instruction"`
	Time        int    `json:"time"`
	TimeUnit    string `json:"timeunit" validate:"timeunit"`
}

//
// Recipe Tag Types
//

type Tag struct {
	Name string `json:"name"`
}

func (v RicettaValidator) validateTimeUnit(i interface{}) error {
	timeUnit := i.(string)
	if timeUnit == "" {
		return fmt.Errorf("Required field")
	} else if !v.Constants.TIME_UNIT_REGEX.MatchString(timeUnit) {
		return fmt.Errorf(timeUnit + " is not a valid unit of time (secs-weeks)")
	} else {
		return nil
	}
}
