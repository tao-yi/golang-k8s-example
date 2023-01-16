package main

import (
	"fmt"
	"golang-web-app/book"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	port       string
	dbHost     string
	dbPort     int
	dbUsername string
	dbPassword string
	dbDatabase string
	db         *sql.DB
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	var err error
	port = os.Getenv("PORT")
	dbHost = os.Getenv("DB_HOST")
	dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	panicOnError(err)
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_DATABASE")
	panicOnError(err)
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbDatabase)
	// open database
	db, err = sql.Open("postgres", psqlconn)
	panicOnError(err)
	panicOnError(db.Ping())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", book.GetBooks(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/books/{id}", book.GetBook(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/books", book.CreateBook(db)).Methods(http.MethodPost)
	r.HandleFunc("/api/books/{id}", book.UpdateBook(db)).Methods(http.MethodPut)
	r.HandleFunc("/api/books/{id}", book.DeleteBook(db)).Methods(http.MethodDelete)

	quit := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		errChan <- http.ListenAndServe(":"+port, r)
	}()

	fmt.Printf("server started on port: %s\n", port)
	select {
	case <-quit:
		fmt.Println("received signal to terminate")
		os.Exit(0)
	case err := <-errChan:
		log.Fatal(err)
	}
}
