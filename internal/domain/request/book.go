package request

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishYear string `json:"publish_year"`
	ISBN        string `json:"isbn"`
}
