package book

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		var book Book
		err := db.QueryRow(`
			SELECT "id","isbn","title","description","created_at" 
			FROM "book"
			WHERE "id"=$1
		`, ID).Scan(&book.ID, &book.Isbn, &book.Title, &book.Description, &book.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(book)
	}
}
