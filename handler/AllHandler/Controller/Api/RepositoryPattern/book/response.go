package book

type BookResponse struct {
	ID			int		`json:"id"`
	Title		string	`json:"title"`
	Age			int		`json:"age"`
	Description string	`json:"description"`
	Price		int		`json:"price"`
	Rating		int		`json:"rating"`
	Discount	int		`json:"discount"`
}