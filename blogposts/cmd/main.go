package main

import (
	"fmt"
	"go-by-tests/blogposts"
	"os"
)

func main() {
	// NOTE: blogs folder is in the main directory
	posts, err := blogposts.NewPostsFromFS(os.DirFS("blogs"))
	if err != nil {
		panic(err)
	}

	fmt.Println(posts)
}
