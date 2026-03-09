package blog

import (
	"errors"
	"fmt"
	"time"
)

// This is metadata of the blog, it will have title, content and createdAt properties. It also have a method to display the content of the blog.

type BlogData struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *BlogData) DisplayContent() {
	fmt.Printf("Blog Title: %s\nBlog Content: %s\n", b.Title, b.Content)
}

func New(title, content string) (*BlogData, error) {
	if len(title) == 0 {
		return nil, errors.New("Title cannot be empty")
	} else if len(content) == 0 {
		return nil, errors.New("Content cannot be empty")
	} else {
		return &BlogData{
			title,
			content,
			time.Now(),
		}, nil
	}

}
