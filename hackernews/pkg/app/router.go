package app

import (
	"net/http"
	"sort"
	"strings"

	"github.com/google/uuid"

	"github.com/gocopper/copper/cerrors"
	"github.com/gocopper/copper/chttp"
	"github.com/gocopper/copper/clogger"
	"github.com/gocopper/examples/hackernews/pkg/posts"
)

type NewRouterParams struct {
	Posts  *posts.Queries
	RW     *chttp.ReaderWriter
	Logger clogger.Logger
}

func NewRouter(p NewRouterParams) *Router {
	return &Router{
		posts:  p.Posts,
		rw:     p.RW,
		logger: p.Logger,
	}
}

type Router struct {
	posts  *posts.Queries
	rw     *chttp.ReaderWriter
	logger clogger.Logger
}

func (ro *Router) Routes() []chttp.Route {
	return []chttp.Route{
		{
			Path:    "/",
			Methods: []string{http.MethodGet},
			Handler: ro.HandleIndexPage,
		},

		{
			Path:    "/submit",
			Methods: []string{http.MethodGet},
			Handler: ro.HandleSubmitPage,
		},

		{
			Path:    "/submit",
			Methods: []string{http.MethodPost},
			Handler: ro.HandleSubmitPost,
		},

		{
			Path:    "/posts/{id}/vote",
			Methods: []string{http.MethodPost},
			Handler: ro.HandleSubmitVote,
		},
	}
}

func (ro *Router) HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	allPosts, err := ro.posts.ListPosts(r.Context())
	if err != nil {
		ro.rw.WriteHTMLError(w, r, cerrors.New(err, "failed to list posts", nil))
		return
	}

	voteCountByPostIDs, err := ro.posts.VoteCountByPostIDs(r.Context(), posts.PostIDs(allPosts))
	if err != nil {
		ro.rw.WriteHTMLError(w, r, cerrors.New(err, "failed to count votes for all posts", nil))
		return
	}

	for i := range allPosts {
		allPosts[i].Votes = voteCountByPostIDs[allPosts[i].ID]
	}

	sort.Slice(allPosts, func(i, j int) bool {
		return allPosts[i].Votes > allPosts[j].Votes
	})

	for i := range allPosts {
		allPosts[i].Rank = int64(i + 1)
	}

	ro.rw.WriteHTML(w, r, chttp.WriteHTMLParams{
		PageTemplate: "index.html",
		Data: map[string][]posts.Post{
			"Posts": allPosts,
		},
	})
}

func (ro *Router) HandleSubmitPage(w http.ResponseWriter, r *http.Request) {
	ro.rw.WriteHTML(w, r, chttp.WriteHTMLParams{
		PageTemplate: "submit.html",
	})
}

func (ro *Router) HandleSubmitPost(w http.ResponseWriter, r *http.Request) {
	var (
		title = strings.TrimSpace(r.PostFormValue("title"))
		url   = strings.TrimSpace(r.PostFormValue("url"))
	)

	if title == "" || url == "" {
		ro.rw.WriteHTMLError(w, r, cerrors.New(nil, "invalid post", map[string]interface{}{
			"form": r.Form,
		}))
		return
	}

	err := ro.posts.SavePost(r.Context(), &posts.Post{
		ID:     uuid.New().String(),
		Title:  title,
		URL:    url,
		Poster: posts.User,
	})
	if err != nil {
		ro.rw.WriteHTMLError(w, r, cerrors.New(err, "failed to save post", map[string]interface{}{
			"form": r.Form,
		}))
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (ro *Router) HandleSubmitVote(w http.ResponseWriter, r *http.Request) {
	var postID = chttp.URLParams(r)["id"]

	err := ro.posts.SaveVote(r.Context(), &posts.Vote{
		ID:     uuid.New().String(),
		PostID: postID,
		User:   posts.User,
	})
	if err != nil {
		ro.rw.WriteHTMLError(w, r, cerrors.New(err, "failed to save vote", map[string]interface{}{
			"postID": postID,
		}))
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
