package book

type PostBookRequest struct {
	Title     string `json:"title" form:"title"`
	Isbn    string `json:"isbn" form:"isbn"`
	Description   string `json:"description" form:"description"`
}

type EditBookRequest struct {
	Title     string `json:"title" form:"title"`
	Isbn    string `json:"isbn" form:"isbn"`
	Description   string `json:"description" form:"description"`
}