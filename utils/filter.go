package utils

import (
	"games/models"
	"math"
	"reflect"
	"time"
)

func Filter(list []models.Restaurant, fieldName string, value interface{}, operator string) (result []models.Restaurant) {
	if fieldName == "-" {
		return list
	}

	for _, v := range list {
		var modelValue, criteriaValue interface{}
		modelValue = reflect.ValueOf(v).FieldByName(fieldName).Interface()
		switch fieldName {
		case "CostBracket":
			modelValue, criteriaValue = modelValue.(int), value.(int)
		case "Cuisine":
			modelValue, criteriaValue = modelValue.(string), value.(string)
		case "IsRecommended":
			modelValue, criteriaValue = modelValue.(bool), value.(bool)
		case "Rating":
			modelValue, criteriaValue = modelValue.(float64), value.(float64)
		case "OnboardedTime":
			modelValue = math.Round(time.Now().Sub(modelValue.(time.Time)).Hours() / 24)
			criteriaValue = value.(float64)
		}
		switch operator {
		case ">=":
			{
				if modelValue.(float64) >= criteriaValue.(float64) {
					result = append(result, v)
				}
			}
		case "<":
			{
				if modelValue.(float64) < criteriaValue.(float64) {
					result = append(result, v)
				}
			}
		default:
			{
				if modelValue == criteriaValue {
					result = append(result, v)
				}
			}
		}
	}
	return result
}
