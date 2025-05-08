package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Slug  string `gorm:"uniqueindex:idx_slug"`
	Likes uint
}

func (p Post) String() string {
	return fmt.Sprintf("Post Title %s, Slug %s", p.Title, p.Slug)
}

var db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func main() {

	// db.AutoMigrate(&Post{})

	// freshPost := createPost("New post title", "New slug")

	fmt.Println(getPost("slug"))
}

func createPost(title string, slug string) Post {
	newPost := Post{Title: title, Slug: slug}
	db.Create(&newPost)
	return newPost
}

func getPost(slug string) Post {
	targetPost := Post{Slug: slug}

	if res := db.First(&targetPost); res.Error != nil {
		panic(res.Error)
	}

	return targetPost
}
