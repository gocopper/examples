package posts

const User = "user1"

type Post struct {
	ID     string `gorm:"primaryKey"`
	Title  string
	URL    string
	Poster string

	Rank  int64 `gorm:"-"`
	Votes int64 `gorm:"-"`
}

type Vote struct {
	ID     string `gorm:"primaryKey"`
	PostID string
	User   string
}
