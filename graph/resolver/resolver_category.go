package graph

import (
	"log"

	"github.com/mskarbe/go-gql-api-server/graph/model"
)

// get all the categories
func (r *Resolver) getCategories() ([]*model.Category, error) {
	sql := `SELECT * FROM category;`
	var categories []*model.Category
	
	rows, err := r.DbSchema.Query(sql)
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
	sql := `SELECT * FROM category where category_id=?;`
	var category model.Category
	err := r.DbSchema.QueryRow(sql, category_id).Scan(&category.ID, &category.Comment)
	if err != nil {
		log.Println(err) 
		return nil, err
	}

	return &category, nil
}
