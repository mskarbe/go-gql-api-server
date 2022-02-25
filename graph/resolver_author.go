package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)

// Insert new author into database
func (r *Resolver) dbInsertAuthor(ctx context.Context, author *model.Author) (error) {
	sql := `INSERT INTO author (author_id, full_name, description, photo_url) VALUES ($1, $2, $3, $4) RETURNING author_id`
	var id string

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	
	err = query.QueryRow(author.ID, author.FullName, author.Description, author.PhotoURL).Scan(&id)
	if err != nil {
  		log.Println(err)
		  return err
	}
	log.Println("New Author record:", id)
	return nil
}

// get all the authors from the database with nested books
func (r *Resolver) getAuthors() ([]*model.Author, error) {
	sql := `SELECT author_id FROM author;`
	var authors []*model.Author

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := query.Query()
	if err != nil {
  		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var author_id string
        if err = rows.Scan(&author_id); err != nil {
			log.Println(err)
		  	return nil, err
		}
		author, err := r.getAuthorByPk(author_id)
		if err != nil {
			log.Println(err)
		  	return nil, err
		}

		authors = append(authors, author)
    }

	return authors, nil
}

// get a single author by id
func (r *Resolver) getAuthorByPk(author_id string) (*model.Author, error) {
	sql := `SELECT * FROM author WHERE author_id=$1::VARCHAR(25);`
	var author model.Author

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	err = query.QueryRow(author_id).Scan(&author.ID, &author.FullName, &author.Description, &author.PhotoURL)
	if err != nil {
		log.Println(err) 
		return nil, err
	}

	author.Books, err = r.getAuthorBooks(author_id)
	if err != nil {
		log.Println(err) 
		return nil, err
	}

	return &author, nil
}

// get a single author's name by id
func (r *Resolver) getAuthorNameByPk(author_id string) (*model.Author, error) {
	sql := `SELECT author_id, full_name FROM author WHERE author_id=$1::VARCHAR(25);`
	var author model.Author

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	err = query.QueryRow(author_id).Scan(&author.ID, &author.FullName)
	if err != nil {
		log.Println(err) 
		return nil, err
	}

	return &author, nil
}

// get all author's books
func (r *Resolver) getAuthorBooks(author_id string) ([]*model.Book, error) {
	sql := `SELECT * FROM book_author WHERE author_id=$1::VARCHAR(25);`
	var books []*model.Book

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := query.Query(sql)
	if err != nil {
  		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var book_id string
		var retrieved_author_id string
		var book *model.Book
        if err = rows.Scan(&book_id, &retrieved_author_id); err != nil {
			log.Println(err)
		  	return nil, err
		}
		// safety check
		if (retrieved_author_id != author_id) {
			return nil, fmt.Errorf("getAuthorBooks: retrieved author_id value does not match requested")
		}
		if book, err = r.getBookByPk(book_id); err != nil {
			log.Println(err)
		  	return nil, err
		}
		books = append(books, book)
    }

	return books, nil
}
