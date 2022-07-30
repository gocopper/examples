package posts

import (
	"context"
	"database/sql"

	"github.com/gocopper/copper/csql"
)

var ErrRecordNotFound = sql.ErrNoRows

func NewQueries(querier csql.Querier) *Queries {
	return &Queries{
		querier: querier,
	}
}

type Queries struct {
	querier csql.Querier
}

func (q *Queries) SavePost(ctx context.Context, post *Post) error {
	const query = "INSERT INTO posts (id, title, url, poster) values (?, ?, ?, ?)"

	_, err := q.querier.Exec(ctx, query,
		post.ID,
		post.Title,
		post.URL,
		post.Poster,
	)
	return err
}

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	const query = "SELECT id, title, url, poster FROM posts"

	var (
		posts []Post
		err   = q.querier.Select(ctx, &posts, query)
	)

	return posts, err
}

func (q *Queries) SaveVote(ctx context.Context, vote *Vote) error {
	const query = "INSERT INTO votes (id, post_id, user) values (?, ?, ?)"

	_, err := q.querier.Exec(ctx, query,
		vote.ID,
		vote.PostID,
		vote.User,
	)
	return err
}

func (q *Queries) VoteCountByPostIDs(ctx context.Context, postIDs []string) (map[string]int64, error) {
	const query = "SELECT post_id, count(*) AS count FROM votes WHERE post_id IN (?) GROUP BY post_id"

	var votes []struct {
		PostID string `db:"post_id"`
		Count  int64
	}

	if len(postIDs) == 0 {
		return map[string]int64{}, nil
	}

	err := q.querier.WithIn().Select(ctx, &votes, query, postIDs)
	if err != nil {
		return nil, err
	}

	voteCountByPostIDs := make(map[string]int64, len(votes))
	for i := range votes {
		voteCountByPostIDs[votes[i].PostID] = votes[i].Count
	}

	return voteCountByPostIDs, nil
}
