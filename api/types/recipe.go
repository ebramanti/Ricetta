package types

import (
	"fmt"
	"unicode/utf8"
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
	PrepTimeUnit string       `json:"preptimeunit" validate:"timeunit"`
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
		return fmt.Errorf(timeUnit + " is not a valid unit of time - [sec(s), min(s), hr(s), day(s), week(s)]")
	} else {
		return nil
	}
}

func (v RicettaValidator) validateTime(i interface{}) error {
	time, ok := i.(int)
	if !ok {
		return fmt.Errorf("Required field")
	} else if time <= 0 {
		return fmt.Errorf("Invalid time unit: %d", time)
	} else {
		return nil
	}
}

func (v RicettaValidator) validateRecipeTitle(i interface{}) error {
	title := i.(string)
	titlelen := utf8.RuneCountInString(title)
	if title == "" {
		return fmt.Errorf("Required field")
	} else if titlelen > v.Constants.MAX_RECIPE_TITLE_LENGTH {
		return fmt.Errorf("max length is %d", v.Constants.MAX_RECIPE_TITLE_LENGTH)
	} else {
		return nil
	}
}

func (v RicettaValidator) validateRecipeNotes(i interface{}) error {
	notes := i.(string)
	noteslen := utf8.RuneCountInString(notes)
	if notes == "" {
		return fmt.Errorf("Required field")
	} else if noteslen > v.Constants.MAX_RECIPE_NOTES_LENGTH {
		return fmt.Errorf("max length is %d", v.Constants.MAX_RECIPE_NOTES_LENGTH)
	} else {
		return nil
	}
}
