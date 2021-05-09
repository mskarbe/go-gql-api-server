package graph

import (
	"database/sql"
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)

// Resolver: holds database connection and all the gql resources
type Resolver struct{
	DbSchema *sql.DB

	books []*model.Book
	authors []*model.Author
	categories []*model.Category
	formats []*model.Format
	format_types []*model.FormatType
}


// initialize the repository: fetch all the existing records from the database
// on the startup of the server 
func (r *Resolver) Initialize() error{
	var err error

	if r.books, err = r.getBooks(); err != nil {
		log.Println(err)
		return err
	}

	if r.authors, err = r.getAuthors(); err != nil {
		log.Println(err)
		return err
	}

	if r.categories, err = r.getCategories(); err != nil {
		log.Println(err)
		return err
	}

	if r.formats, err = r.getFormats(); err != nil {
		log.Println(err)
		return err
	}

	if r.format_types, err = r.getFormatTypes(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
