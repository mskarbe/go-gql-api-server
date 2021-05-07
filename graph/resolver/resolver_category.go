package graph

import (
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)

// get all the categories
func (r *Resolver) getCategories() ([]*model.Category, error) {
	sql := `SELECT * FROM category;`
	var categories []*model.Category

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
        var category model.Category
        if err = rows.Scan(&category.ID, &category.Comment); err != nil {
			log.Println(err)
		  	return nil, err
		}
		categories = append(categories, &category)
    }

	return categories, nil
}

// get a single category by id
func (r *Resolver) getCategoryByPk(category_id string) (*model.Category, error) {
	sql := `SELECT * FROM category WHERE category_id=$1;`
	var category model.Category

	query, err := r.DbSchema.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	err = query.QueryRow(category_id).Scan(&category.ID, &category.Comment)
	if err != nil {
		log.Println(err) 
		return nil, err
	}

	return &category, nil
}
