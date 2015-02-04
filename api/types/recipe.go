package types

//
// Recipe Creation Types
//

type Recipe struct {
	Title       string
	Notes       string
	Ingredients []Ingredient
	// Time is represented in minutes
	CookTime     int
	CookTimeUnit string
	PrepTime     int
	PrepTimeUnit string
	Steps        []Step
	Tags         []Tag
	Public       bool
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
