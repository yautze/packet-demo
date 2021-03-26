package go_request

// TestBody -
type TestBody struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
