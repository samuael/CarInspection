package listing

// Post defines the properties of a Post to be listed
type Inspection struct {
	ID       uint   `json:"id"`
	AuthorID uint   `json:"author_id"`
	Content  string `json:"content"`
}
