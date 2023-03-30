package main

import (
	"fmt"
	"games/constants"
	"games/models"
	"games/utils"
	"sort"
	"sync"
	"time"
)

func main() {
	// Cuisines    []CuisineTracking
	// CostBracket []CostTracking
	user := models.User{
		Cuisines: []models.CuisineTracking{
			{"Chinese", 20},
			{"NorthIndian", 5},
			{"Biryani", 3},
		},
		CostBracket: []models.CostTracking{
			{1, 4},
			{4, 4},
			{5, 6},
		},
	}
	availableRestaurants := []models.Restaurant{
		{RestaurantID: "1", Cuisine: "Chinese", CostBracket: 5, Rating: 2, IsRecommended: false, OnboardedTime: time.Now().AddDate(0, 0, -2)},
		{RestaurantID: "2", Cuisine: "Chinese", CostBracket: 1, Rating: 4, IsRecommended: true, OnboardedTime: time.Now().AddDate(0, 0, -2)},
		{RestaurantID: "3", Cuisine: "NorthIndian", CostBracket: 3, Rating: 4, IsRecommended: true, OnboardedTime: time.Now().AddDate(0, 0, -2)},
		{RestaurantID: "4", Cuisine: "Chinese", CostBracket: 4, Rating: 4, IsRecommended: true, OnboardedTime: time.Now().AddDate(0, 0, -2)},
		{RestaurantID: "5", Cuisine: "Biryani", CostBracket: 5, Rating: 4, IsRecommended: true, OnboardedTime: time.Now().AddDate(0, 0, -2)},
	}

	for _, v := range GetUserRecommendations(user, availableRestaurants) {
		fmt.Println(v)
	}
}

func GetUserRecommendations(user models.User, availableRestaurants []models.Restaurant) []string {
	var wg sync.WaitGroup
	ch := make(chan models.RespRestaurants)
	receiceRestaurants := make([][]models.Restaurant, len(constants.RulesForRestaurantRecommendation))
	respRestaurantIDs := make([]string, 0)
	visitedRestaurantMap := make(map[string]bool)
	userPrimaryCuisine, userSecondaryCuisines := getPrimarySecondaryCuisine(user)
	user.PrimaryCuisine = userPrimaryCuisine
	user.SecondaryCuisine = userSecondaryCuisines

	userPrimaryCost, userSecondaryCost := getPrimarySecondaryCost(user)
	user.PrimaryCostBracket = userPrimaryCost
	user.SecondaryCostBracket = userSecondaryCost

	fmt.Println("primary cuisine - ", user.PrimaryCuisine)
	fmt.Println("secondary cuisine - ", user.SecondaryCuisine)
	fmt.Println("primary cost - ", user.PrimaryCostBracket)
	fmt.Println("secondary cost - ", user.SecondaryCostBracket)

	for _, v := range constants.RulesForRestaurantRecommendation {
		if v.Active {
			wg.Add(1)
			go executeRule(user, v, availableRestaurants, ch)
			go func() {
				defer wg.Done()
				res := <-ch
				ruleSetID := res.RuleSetID
				receiceRestaurants[ruleSetID-1] = res.Restaurants
			}()
		}
	}

	wg.Wait()
	close(ch)

	for i, _ := range receiceRestaurants {
		if constants.RulesForRestaurantRecommendation[i].DependantRuleID > 0 && len(receiceRestaurants[i]) > 0 {
			continue
		}

		for _, restaurant := range receiceRestaurants[i] {
			if len(respRestaurantIDs) >= 100 {
				respRestaurantIDs = respRestaurantIDs[:100]
				break
			}
			if _, exists := visitedRestaurantMap[restaurant.RestaurantID]; !exists {
				respRestaurantIDs = append(respRestaurantIDs, restaurant.RestaurantID)
				visitedRestaurantMap[restaurant.RestaurantID] = true
			}
		}
	}

	return respRestaurantIDs
}

func getPrimarySecondaryCuisine(user models.User) (primary models.CuisineTracking, secondary []models.CuisineTracking) {
	if len(user.Cuisines) == 0 {
		return
	}
	sort.Slice(user.Cuisines, func(i, j int) bool {
		return user.Cuisines[i].NoOfOrders > user.Cuisines[j].NoOfOrders
	})
	primary = user.Cuisines[0]
	secondaryCuisines := user.Cuisines[1:len(user.Cuisines)]
	if len(secondaryCuisines) > 0 {
		secondary = append(secondary, secondaryCuisines[0])
		if len(secondaryCuisines) > 1 {
			secondary = append(secondary, secondaryCuisines[1])
		}
	}
	return
}

func getPrimarySecondaryCost(user models.User) (primary models.CostTracking, secondary []models.CostTracking) {
	if len(user.CostBracket) == 0 {
		return
	}
	sort.Slice(user.CostBracket, func(i, j int) bool {
		return user.CostBracket[i].NoOfOrders > user.CostBracket[j].NoOfOrders
	})
	primary = user.CostBracket[0]
	secondaryCosts := user.CostBracket[1:len(user.CostBracket)]
	if len(secondaryCosts) > 0 {
		secondary = append(secondary, secondaryCosts[0])
		if len(secondaryCosts) > 1 {
			secondary = append(secondary, secondaryCosts[1])
		}
	}
	return
}

func handleCuisineFilter(user models.User, availableRestaurants []models.Restaurant, rule models.Rule) []models.Restaurant {
	if rule.Value == "primary" {
		availableRestaurants = utils.Filter(availableRestaurants, rule.Field, user.PrimaryCuisine.CuisineType, "")
	} else if rule.Value == "secondary" {
		res := []models.Restaurant{}
		for _, cuisine := range user.SecondaryCuisine {
			res = append(res, utils.Filter(availableRestaurants, rule.Field, cuisine.CuisineType, "")...)
		}
		availableRestaurants = res
	}
	return availableRestaurants
}

func handleCostBracketFilter(user models.User, availableRestaurants []models.Restaurant, rule models.Rule) []models.Restaurant {
	if rule.Value == "primary" {
		availableRestaurants = utils.Filter(availableRestaurants, rule.Field, int(user.PrimaryCostBracket.Type), "")
	} else if rule.Value == "secondary" {
		res := []models.Restaurant{}
		for _, cost := range user.SecondaryCostBracket {
			res = append(res, utils.Filter(availableRestaurants, rule.Field, int(cost.Type), "")...)
		}
		availableRestaurants = res
	}
	return availableRestaurants
}

func executeRule(user models.User, ruleset models.Ruleset, availableRestaurants []models.Restaurant, ch chan<- models.RespRestaurants) {
	defer func() {
		ch <- models.RespRestaurants{RuleSetID: ruleset.RuleSetID, Restaurants: availableRestaurants}
	}()

	for _, v := range ruleset.Rules {
		switch v.Field {
		case "Cuisine":
			availableRestaurants = handleCuisineFilter(user, availableRestaurants, v)
		case "CostBracket":
			availableRestaurants = handleCostBracketFilter(user, availableRestaurants, v)
		default:
			availableRestaurants = utils.Filter(availableRestaurants, v.Field, v.Value, v.Operator)
		}
	}
	if ruleset.Sort.Field != "" {
		sort.Slice(availableRestaurants, func(i, j int) bool {
			return availableRestaurants[i].Rating > availableRestaurants[j].Rating
		})
	}
	if ruleset.Limit > 0 && len(availableRestaurants) > ruleset.Limit {
		availableRestaurants = availableRestaurants[0:ruleset.Limit]
	}
}
