package book

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		stmt := `update "book" set "description"=$1 where "id"=$2`
		_, err := db.Exec(stmt, book.Description, ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "success",
		})
	}
}
