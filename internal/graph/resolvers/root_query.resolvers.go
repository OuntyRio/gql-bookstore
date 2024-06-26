package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/ountyrio/gql-bookstore/internal/graph/generated"
	gqlmodel "github.com/ountyrio/gql-bookstore/internal/graph/model"
)

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) (*gqlmodel.AuthorsQueryNs, error) {
	return &gqlmodel.AuthorsQueryNs{}, nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) (*gqlmodel.BooksQueryNs, error) {
	return &gqlmodel.BooksQueryNs{}, nil
}

// Genres is the resolver for the genres field.
func (r *queryResolver) Genres(ctx context.Context) (*gqlmodel.GenresQueryNs, error) {
	return &gqlmodel.GenresQueryNs{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
