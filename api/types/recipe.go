package types

//
// Recipe Creation Types
//

type Recipe struct {
	Title        string
	Notes        string
	Ingredients  []Ingredient
	CookTime     int
	CookTimeUnit string
	PrepTime     int
	PrepTimeUnit string
	Steps        []Step
	Tags         []Tag
	Public       bool
	CookTimeUnit string       `json:"cooktimeunit" validate:"timeunit"`
	PrepTimeUnit string       `json:"prepunit" validate:"timeunit"`
}

//
// Recipe Component Types
//

type Ingredient struct {
	Name   string
	Amount string
	URL    string
}

type Step struct {
	Instruction string
	Time        int
	TimeUnit    string
}

//
// Recipe Tag Types
//

type Tag struct {
	Name string
}

func (v RicettaValidator) validateTimeUnit(i interface{}) error {
	timeUnit := i.(string)
	if timeUnit == "" {
		return fmt.Errorf("Required field")
	} else if !v.Constants.TIME_UNIT_REGEX.MatchString() {
		return fmt.Errorf(handle + " is not a valid unit of time (secs-weeks)")
	} else {
		return nil
	}
}
