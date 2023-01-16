package book

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		stmt := `insert into "book"("isbn", "title", "description") values($1, $2, $3)`
		_, err := db.Exec(stmt, book.Isbn, book.Title, book.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("success"))
	}
}
