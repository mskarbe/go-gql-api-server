package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)

// Insert new book into database
func (r *Resolver) DbInsertBook(ctx context.Context, book *model.Book) (error) {
	sql := `INSERT INTO book (book_id, title, year, publisher, description, photo_url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING book_id`
	var id string
	
	err := r.DbSchema.QueryRow(sql, book.ID, book.Title, book.Year, book.Publisher, book.Description, book.CoverURL).Scan(&id)
	if err != nil {
  		log.Println(err)
		  return err
	}
	log.Println("New book record:", id)
	return nil
}

// get all the book records from database with nested categories, formats and authors
func (r *Resolver) getBooks() ([]*model.Book, error) {
	sql := `SELECT book_id FROM book;`
	var books []*model.Book
	
	rows, err := r.DbSchema.Query(sql)
	if err != nil {
  		log.Println(err)
		return nil, err
	}

	for rows.Next() {
        var book_id string
        if err = rows.Scan(&book_id); err != nil {
			log.Println(err)
		  	return nil, err
		}
		book, err := r.getBookByPk(book_id)
		if err != nil {
			log.Println(err)
		  	return nil, err
		}

		books = append(books, book)
    }

	return books, nil
}

// get a single author by id
func (r *Resolver) getBookByPk(book_id string) (*model.Book, error) {
	sql := `SELECT * FROM book WHERE book_id=?;`

	var book model.Book
    err := r.DbSchema.QueryRow(sql, book_id).Scan(&book.ID, &book.Title, &book.Year, &book.Publisher, &book.Description, &book.CoverURL)
	if err != nil {
		log.Println(err)
	  	return nil, err
	}
	// retrieve book's formats
	formats, err := r.getBookFormats(book.ID)
	if err != nil {
		log.Println(err)
	  	return nil, err
	}
	book.Formats = formats 
	// retrieve book's authors
	authors, err := r.getBookAuthors(book.ID)
	if err != nil {
		log.Println(err)
	  	return nil, err
	}
	book.Authors = authors 
	// retrieve book's categories 
	categories, err := r.getBookCategories(book.ID)
	if err != nil {
		log.Println(err)
	  	return nil, err
	}
	book.Categories = categories 

	return &book, nil
}

func (r *Resolver) getBookAuthors(book_id string) ([]*model.Author, error) {
	sql := `SELECT * FROM book_author WHERE book_id=?;`
	var authors []*model.Author
	
	rows, err := r.DbSchema.Query(sql)
	if err != nil {
  		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var author_id string
		var retrieved_book_id string
		var author *model.Author
        if err = rows.Scan(&retrieved_book_id, &author_id); err != nil {
			log.Println(err)
		  	return nil, err
		}
		// safety check
		if (retrieved_book_id != book_id) {
			return nil, fmt.Errorf("getBookAuthors: retrieved book_id value does not match requested")
		}
		if author, err = r.getAuthorByPk(author_id); err != nil {
			log.Println(err)
		  	return nil, err
		}
		authors = append(authors, author)
    }

	return authors, nil
}

// gets all categories of a single book
func (r *Resolver) getBookCategories(book_id string) ([]*model.Category, error) {
	sql := `SELECT * FROM book_category WHERE book_id=?;`
	var category_id string
	var retrieved_book_id string
	var categories []*model.Category
	
	rows, err := r.DbSchema.Query(sql, book_id)
	if err != nil {
  		log.Println(err)
		return nil, err
	}

	for rows.Next() {
        if err = rows.Scan(&retrieved_book_id, &category_id); err != nil {
			log.Println(err)
		  	return nil, err
		}
		// safety check
		if (retrieved_book_id != book_id) {
			return nil, fmt.Errorf("getBookCategories: retrieved book_id value does not match requested")
		}
		categories = append(categories, &model.Category{ID: category_id})
    }

	return categories, nil
}


// get all formats of a single book
func (r *Resolver) getBookFormats(book_id string) ([]*model.Format, error) {
	sql := `SELECT * FROM format WHERE book_id=?;`
	var book_formats []*model.Format
	
	rows, err := r.DbSchema.Query(sql, book_id)
	if err != nil {
  		log.Println(err)
		return nil, err
	}

	for rows.Next() {
        var format model.Format
        if err = rows.Scan(&format.ID, &format.Book.ID, &format.Type.ID, &format.Price, &format.Supply); err != nil {
			log.Println(err)
		  	return nil, err
		}
		// safety check
		if (format.Book.ID != book_id) {
			return nil, fmt.Errorf("getBookFormats: retrieved book_id value does not match requested")
		}

		book_formats = append(book_formats, &format)
    }

	return book_formats, nil
}
