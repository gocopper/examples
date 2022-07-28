package posts

const User = "user1"

type Post struct {
	ID     string
	Title  string
	URL    string
	Poster string

	Rank  int64
	Votes int64
}

type Vote struct {
	ID     string
	PostID string
	User   string
}
