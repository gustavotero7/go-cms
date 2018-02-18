package cms

import (
	"net/http"
	"strings"
	"time"
)

// ServePage _
func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")
	if path == "" {
		pages, err := GetPages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Tmpl.ExecuteTemplate(w, "pages", pages)
		return
	}

	page, err := GetPage(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	Tmpl.ExecuteTemplate(w, "page", page)

}

// ServePost _
func ServePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")
	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Post{
		Title:   strings.ToTitle(path),
		Content: "Here is my page",
	}

	Tmpl.ExecuteTemplate(w, "post", p)
}

// HandleNew _
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			p := &Page{
				Title:   title,
				Content: content,
			}

			_, err := CreatePage(p)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			Tmpl.ExecuteTemplate(w, "page", p)
		} else if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "page", &Post{
				Title:   title,
				Content: content,
			})
		}
	default:
		http.Error(w, "Method not supported "+r.Method, http.StatusMethodNotAllowed)
	}
}

// ServeIndex _
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title:   "",
		Content: "",
		Posts: []*Post{
			{
				Title:         "Hello world",
				Content:       "Hello world! this is my first templating excercice",
				DatePublished: time.Now(),
				Comments: []*Comment{
					{
						Author:        "Some guy",
						Comment:       "Yeah, this is f*king great!",
						DatePublished: time.Now(),
					},
				},
			},
			{
				Title:         "2nd post",
				Content:       "Another post in this brand new cms",
				DatePublished: time.Now(),
			},
			{
				Title:         "Yeah!",
				Content:       "Yep, another useless post with nothig to see",
				DatePublished: time.Now(),
				Comments: []*Comment{
					{
						Author:        "Guy 2",
						Comment:       "Humm, keep going, amazin this are comming!",
						DatePublished: time.Now(),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}
