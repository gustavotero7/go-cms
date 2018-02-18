package cms

import (
	"html/template"
	"time"
)

// Tmpl Template
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

// Page struct
type Page struct {
	ID      int
	Title   string
	Content string
	Posts   []*Post
}

// Post _
type Post struct {
	ID            int
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}

// Comment _
type Comment struct {
	ID            int
	Author        string
	Comment       string
	DatePublished time.Time
}
