package models

import "time"

type (
	Restaurant struct {
		RestaurantID  string
		Cuisine       string // cuisine data type
		CostBracket   int
		Rating        float64
		IsRecommended bool
		OnboardedTime time.Time
	}

	CuisineTracking struct {
		CuisineType string
		NoOfOrders  int
	}

	CostTracking struct {
		Type       int
		NoOfOrders int
	}

	User struct {
		Cuisines             []CuisineTracking
		CostBracket          []CostTracking
		PrimaryCuisine       CuisineTracking
		SecondaryCuisine     []CuisineTracking
		PrimaryCostBracket   CostTracking
		SecondaryCostBracket []CostTracking
	}
)

type Sort struct {
	Field string
	Type  string
}
type Ruleset struct {
	RuleSetID       int
	Rules           []Rule
	Active          bool
	Sort            Sort
	Limit           int
	SubRules        []Rule
	DependantRuleID int
}

type Rule struct {
	Field     string
	Operator  string
	Value     interface{}
	ValueType string
}

type RespRestaurants struct {
	RuleSetID   int
	Restaurants []Restaurant
}
