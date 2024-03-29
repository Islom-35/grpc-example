package domain

type Data struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}


type PostsResponse struct {
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Pages int `json:"pages"`
			Page  int `json:"page"`
			Limit int `json:"limit"`
		} `json:"pagination"`
	} `json:"meta"`
	Data []Data `json:"data"`
}
