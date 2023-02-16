package mocks

import "github.com/rominirani/crud-mockapi/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
	{
		Id:     2,
		Title:  "Golang",
		Author: "Gopher Guide",
		Desc:   "A tour of Go",
	},
	{
		Id:     3,
		Title:  "Golang",
		Author: "Expert Gopher",
		Desc:   "Advanced Go",
	},
}
