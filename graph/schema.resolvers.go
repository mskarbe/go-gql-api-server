package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mskarbe/go-gql-api-server/graph/generated"
	"github.com/mskarbe/go-gql-api-server/graph/model"
)

func (r *mutationResolver) InsertBook(ctx context.Context, title string, year *int, description *string, coverURL *string, authors []string) (*model.Book, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *mutationResolver) InsertCategory(ctx context.Context, id string, comment *string) (*model.Category, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *mutationResolver) InsertFormat(ctx context.Context, book string, price float64, typeArg string, supply int) (*model.Format, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *mutationResolver) InsertFormatType(ctx context.Context, id string, comment *string) (*model.FormatType, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *mutationResolver) InsertAuthor(ctx context.Context, fullName string, description *string, photoURL *string) (*model.Author, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) BookByPk(ctx context.Context, id *string) (*model.Book, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) BooksByFormat(ctx context.Context, format *string) ([]*model.Book, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) BooksByCategory(ctx context.Context, category *string) ([]*model.Book, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) CategoryByPk(ctx context.Context, id *string) (*model.Category, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) FormatByPk(ctx context.Context, id *string) (*model.Format, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) Formats(ctx context.Context) ([]*model.Format, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) FormatTypeByPk(ctx context.Context, id *string) (*model.FormatType, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) FormatTypes(ctx context.Context) ([]*model.FormatType, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) AuthorByPk(ctx context.Context, id *string) (*model.Author, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	fmt.Println("not implemented")
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
