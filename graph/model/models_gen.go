// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID          string  `json:"id"`
	FullName    string  `json:"full_name"`
	Description *string `json:"description"`
	PhotoURL    *string `json:"photo_url"`
	Books       []*Book `json:"books"`
}

type Book struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Year        *int        `json:"year"`
	Publisher   *string     `json:"publisher"`
	Description *string     `json:"description"`
	CoverURL    *string     `json:"cover_url"`
	Authors     []*Author   `json:"authors"`
	Formats     []*Format   `json:"formats"`
	Categories  []*Category `json:"categories"`
}

type Category struct {
	ID      string  `json:"id"`
	Comment *string `json:"comment"`
}

type Format struct {
	ID     string      `json:"id"`
	Book   *Book       `json:"book"`
	Price  float64     `json:"price"`
	Type   *FormatType `json:"type"`
	Supply int         `json:"supply"`
}

type FormatType struct {
	ID      string  `json:"id"`
	Comment *string `json:"comment"`
}
