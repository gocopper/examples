package posts

import (
	"context"

	"github.com/gocopper/copper/cerrors"

	"github.com/gocopper/copper/csql"
	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type Repo struct {
	db *gorm.DB
}

func (r *Repo) SavePost(ctx context.Context, post *Post) error {
	return csql.GetConn(ctx, r.db).Save(post).Error
}

func (r *Repo) ListPosts(ctx context.Context) ([]Post, error) {
	var posts []Post

	err := csql.GetConn(ctx, r.db).
		Find(&posts).
		Error
	if err != nil {
		return nil, cerrors.New(err, "failed to scan posts", nil)
	}

	return posts, nil
}

func (r *Repo) SaveVote(ctx context.Context, vote *Vote) error {
	return csql.GetConn(ctx, r.db).Save(vote).Error
}

func (r *Repo) VoteCountByPostIDs(ctx context.Context, postIDs []string) (map[string]int64, error) {
	var votes []struct {
		PostID string
		Count  int64
	}

	err := csql.GetConn(ctx, r.db).
		Raw("select post_id, count(*) as count from votes where post_id in (?) group by post_id", postIDs).
		Scan(&votes).
		Error
	if err != nil {
		return nil, cerrors.New(err, "failed to query vote counts", map[string]interface{}{
			"postIDs": postIDs,
		})
	}

	voteCountByPostIDs := make(map[string]int64, len(votes))
	for i := range votes {
		voteCountByPostIDs[votes[i].PostID] = votes[i].Count
	}

	return voteCountByPostIDs, nil
}