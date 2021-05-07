package graph

import (
	"context"
	"database/sql"
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)

type Resolver struct{
	DbSchema *sql.DB

	books []*model.Book
	authors []*model.Author
	categories []*model.Category
	formats []*model.Format
	format_types []*model.FormatType
}

func (r *Resolver) DbInsertBook(ctx context.Context, book *model.Book) (error) {
	
	sql := `INSERT INTO book (book_id, year, publisher, description, photo_url) VALUES ($1, $2, $3, $4, $5) RETURNING book_id`
	var id string
	
	err := r.DbSchema.QueryRow(sql, book.ID, book.Year, book.Publisher, book.Description, book.CoverURL).Scan(&id)
	if err != nil {
  		log.Fatal(err)
		  return err
	}
	log.Println("New book record:", id)
	return nil
}

func (r *Resolver) DbInsertAuthor(ctx context.Context, author *model.Author) (error) {
	
	sql := `INSERT INTO author (author_id, full_name, description, photo_url) VALUES ($1, $2, $3, $4) RETURNING author_id`
	var id string
	
	err := r.DbSchema.QueryRow(sql, author.ID, author.FullName, author.Description, author.PhotoURL).Scan(&id)
	if err != nil {
  		log.Fatal(err)
		  return err
	}
	log.Println("New Author record:", id)
	return nil
}
