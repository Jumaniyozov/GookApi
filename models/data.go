package models

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID:            1,
		Title:         "Lord of The Rings. Vol. 1",
		YearPublished: 1978,
		Author: Author{
			Name:     "J.R",
			LastName: "Tolkin",
			BornYear: 1892,
		},
	}

	DB = append(DB, book1)
}

func FindBookById(id int) (Book, bool) {
	var book Book
	var found bool

	for _, val := range DB {
		if val.ID == id {
			book = val
			found = true
			break
		}
	}
	return book, found
}
