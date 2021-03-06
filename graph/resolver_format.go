package graph

import (
	"context"
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)

// Insert new format into database
func (r *Resolver) dbInsertFormat(ctx context.Context, format *model.Format) (error) {
	sql := `INSERT INTO format (format_id, book_id, format_type, price, supply) VALUES ($1, $2, $3, $4, $5) RETURNING author_id`
	var id string

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	
	err = query.QueryRow(format.ID, format.Book.ID, format.Type.ID, format.Price, format.Supply).Scan(&id)
	if err != nil {
  		log.Println(err)
		  return err
	}
	log.Println("New Format record:", id)
	return nil
}

// get all formats
func (r *Resolver) getFormats() ([]*model.Format, error) {
	sql := `SELECT * FROM format;`
	var formats []*model.Format

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
        var format model.Format
		var book_id string
		var format_type string
        if err = rows.Scan(&format.ID, &book_id, &format_type, &format.Price, &format.Supply); err != nil {
			log.Println(err)
		  	return nil, err
		}
		format.Book, err = r.getBookByPk(book_id)
		if err != nil {
			log.Println(err)
		  	return nil, err
		}

		formats = append(formats, &format)
    }

	return formats, nil
}

// get a single format by id
func (r *Resolver) getFormatByPk(format_id string) (*model.Format, error) {
	sql := `SELECT * FROM format WHERE format_id=$1::VARCHAR(20);`
	var format model.Format

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	err = query.QueryRow(format_id).Scan(&format.ID, &format.Book.ID, &format.Type, &format.Price, &format.Supply)
	if err != nil {
		log.Println(err) 
		return nil, err
	}

	return &format, nil
}

// get format types
func (r *Resolver) getFormatTypes() ([]*model.FormatType, error) {
	sql := `SELECT * FROM format_type;`
	var format_types []*model.FormatType

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
        var format_type model.FormatType
        if err = rows.Scan(&format_type.ID, &format_type.Comment); err != nil {
			log.Println(err)
		  	return nil, err
		}
		format_types = append(format_types, &format_type)
    }

	return format_types, nil
}
