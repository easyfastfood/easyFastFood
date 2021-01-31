package graphql

import (
	"context"

	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
)

type queryResolver struct{}

func NewQueryResolver() gqlgen.QueryResolver {
	return new(queryResolver)
}

func (q queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	return &models.User{
		ID:          id,
		FirstName:   "massoni",
		LastName:    "mestre dos ratos",
		PhoneNumber: 9125912571,
		Email:       "ratão@gmail.com",
	}, nil
}
