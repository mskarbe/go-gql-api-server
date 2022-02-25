package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mskarbe/go-gql-api-server/graph/generated"
	"github.com/mskarbe/go-gql-api-server/graph/model"
	cuid "gopkg.in/lucsky/cuid.v1"
)

func (r *mutationResolver) InsertBook(ctx context.Context, title string, year *int, publisher *string, description *string, coverURL *string, authors []*string, categories []*string) (*model.Book, error) {
	var (
		book_authors    []*model.Author
		book_categories []*model.Category
	)
	book_id := cuid.New()

	// assign parameters
	book := &model.Book{
		ID:          book_id,
		Title:       title,
		Year:        year,
		Publisher:   publisher,
		Description: description,
		CoverURL:    coverURL,
	}

	// resolver call to db
	err := r.dbInsertBook(ctx, book)
	if err != nil {
		return nil, err
	}

	// retrieve authors records
	for _, req := range authors {
		found := false
		for _, a := range r.authors {
			if a.ID == *req {
				found = true
				//insert book_author record
				err := r.dbInsertBookAuthor(ctx, book_id, a.ID)
				if err != nil {
					return nil, err
				}
				book_authors = append(book_authors, a)
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("author [id: %s] does not exist", *req)
		}
	}

	// retrieve categories records
	for _, req := range categories {
		found := false
		for _, c := range r.categories {
			if c.ID == *req {
				found = true
				//insert book_category record
				err := r.dbInsertBookCategory(ctx, book_id, c.ID)
				if err != nil {
					return nil, err
				}
				book_categories = append(book_categories, c)
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("category [id: %s] does not exist", *req)
		}
	}

	book.Authors = book_authors
	book.Categories = book_categories

	// append to repository
	r.books = append(r.books, book)
	return book, nil
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
	author := &model.Author{
		ID:          cuid.New(),
		FullName:    fullName,
		Description: description,
		PhotoURL:    photoURL,
	}

	err := r.dbInsertAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	r.authors = append(r.authors, author)
	return author, nil
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
	//r.getBooks()
	return r.books, nil
}

func (r *queryResolver) CategoryByPk(ctx context.Context, id *string) (*model.Category, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.categories, nil
}

func (r *queryResolver) FormatByPk(ctx context.Context, id *string) (*model.Format, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) Formats(ctx context.Context) ([]*model.Format, error) {
	return r.formats, nil
}

func (r *queryResolver) FormatTypeByPk(ctx context.Context, id *string) (*model.FormatType, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) FormatTypes(ctx context.Context) ([]*model.FormatType, error) {
	return r.format_types, nil
}

func (r *queryResolver) AuthorByPk(ctx context.Context, id *string) (*model.Author, error) {
	fmt.Println("not implemented")
	return nil, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return r.authors, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
