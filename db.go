package cms

import (
	"database/sql"
	// postgres driver
	_ "github.com/lib/pq"
)

var store = newDB()

// PgStore _
type PgStore struct {
	DB *sql.DB
}

func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=goprojects password='123' dbname=goprojects sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &PgStore{
		DB: db,
	}
}

// GetPage _
func GetPage(id string) (*Page, error) {
	var p Page
	err := store.DB.QueryRow("SELECT * FROM pages WHERE id=$1", id).Scan(&p.ID, &p.Title, &p.Content)
	return &p, err
}

// GetPages _
func GetPages() ([]*Page, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pages := []*Page{}
	for rows.Next() {
		var p Page
		err = rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &p)
	}
	return pages, nil
}

// CreatePage _
func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO pages (title,content) VALUES($1, $2) RETURNING id", p.Title, p.Content).Scan(&id)
	return id, err
}
