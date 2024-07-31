package blogposts_test

import (
	"errors"
	"go-by-tests/blogposts"

	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (f StubFailingFS) Open(name string) (fs.File, error) {
	return nil, fs.ErrNotExist
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: World
Description: Description 1
Tags: tdd, go
------
Hi
Go`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
-------
Hi
Rust
Lang`
	)
	t.Run("reads and stored new blogpost", func(t *testing.T) {
		files := fstest.MapFS{
			"1st.md": {Data: []byte(firstBody)},
			"2nd.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(files)
		expectedFirstPost := blogposts.Post{
			Title:       "World",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hi
Go`,
		}
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(files) {
			t.Errorf("got %d posts, expected %d", len(posts), len(files))
		}

		assertPostEqual(t, posts[0], expectedFirstPost)
	})

	t.Run("errors when opening blogpost", func(t *testing.T) {
		failingFS := StubFailingFS{}
		_, err := blogposts.NewPostsFromFS(failingFS)
		if err == nil {
			t.Fatal("expected error when opening blogpost")
		}
		if !errors.Is(err, fs.ErrNotExist) {
			t.Errorf("got %v, wanted %v", err, fs.ErrNotExist)
		}
	})
}

func assertPostEqual(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
