package constants

import "games/models"

var RulesForRestaurantRecommendation = []models.Ruleset{
	{
		RuleSetID: 1,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "primary"},
			{Field: "CostBracket", Value: "primary"},
			{Field: "IsRecommended", Value: true},
		},
		Active: true,
	},
	{
		RuleSetID:       2,
		DependantRuleID: 1,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "primary"},
			{Field: "CostBracket", Value: "secondary"},
			{Field: "IsRecommended", Value: true},
		},
		Active: true,
	},
	{
		RuleSetID:       3,
		DependantRuleID: 1,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "secondary"},
			{Field: "CostBracket", Value: "primary"},
			{Field: "IsRecommended", Value: true},
		},
		Active: true,
	},
	{
		RuleSetID: 2,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "primary"},
			{Field: "CostBracket", Value: "primary"},
			{Field: "Rating", Value: 4.0, Operator: ">="},
		},
		Active: true,
	},
	{
		RuleSetID: 3,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "primary"},
			{Field: "CostBracket", Value: "secondary"},
			{Field: "Rating", Value: 4.5, Operator: ">="},
		},
		Active: true,
	},
	{
		RuleSetID: 4,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "secondary"},
			{Field: "CostBracket", Value: "primary"},
			{Field: "Rating", Value: 4.5, Operator: ">="},
		},
		Active: true,
	},
	{
		RuleSetID: 5,
		Rules: []models.Rule{
			{Field: "OnboardedTime", Value: 2.0, Operator: "<="},
		},
		Active: true,
		Sort:   models.Sort{Field: "Rating", Type: "desc"},
		Limit:  4,
	},
	{
		RuleSetID: 6,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "primary"},
			{Field: "CostBracket", Value: "primary"},
			{Field: "Rating", Value: 4.0, Operator: "<"},
		},
		Active: true,
	},
	{
		RuleSetID: 7,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "primary"},
			{Field: "CostBracket", Value: "secondary"},
			{Field: "Rating", Value: 4.5, Operator: "<"},
		},
		Active: true,
	},
	{
		RuleSetID: 8,
		Rules: []models.Rule{
			{Field: "Cuisine", Value: "secondary"},
			{Field: "CostBracket", Value: "primary"},
			{Field: "Rating", Value: 4.5, Operator: "<"},
		},
		Active: true,
	},
	{
		RuleSetID: 9,
		Rules: []models.Rule{
			{Field: "-"},
		},
		Active: true,
		Limit:  100,
	},
}
