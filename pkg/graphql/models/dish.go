package models

import "easyfood/pkg/entity"

func NewDish(dishes ...*entity.Dish) []*Dish {
	result := make([]*Dish, 0)

	for _, dish := range dishes {
		result = append(result, &Dish{
			ID:       dish.Id,
			Name:     dish.Name,
			Price:    dish.Price,
			CookTime: dish.CookTime,
		})
	}

	return result
}
