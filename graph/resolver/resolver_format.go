package graph

import (
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)


func (r *Resolver) getFormats() ([]*model.Format, error) {
	sql := `SELECT * FROM format;`
	var formats []*model.Format
	
	rows, err := r.DbSchema.Query(sql)
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

func (r *Resolver) getFormatTypes() ([]*model.FormatType, error) {
	sql := `SELECT * FROM format_type;`
	var format_types []*model.FormatType
	
	rows, err := r.DbSchema.Query(sql)
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
