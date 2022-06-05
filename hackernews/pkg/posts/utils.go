package posts

func PostIDs(posts []Post) []string {
	ids := make([]string, len(posts))
	for i := range posts {
		ids[i] = posts[i].ID
	}
	return ids
}
