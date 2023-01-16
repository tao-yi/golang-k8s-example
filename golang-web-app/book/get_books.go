package book

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var books []*Book
		rows, err := db.Query(`
			SELECT "id","isbn","title","description","created_at" 
			FROM "book"
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var book Book
			err = rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Description, &book.CreatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			books = append(books, &book)
		}

		_ = json.NewEncoder(w).Encode(books)
	}
}
