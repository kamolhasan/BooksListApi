package api

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

//-----------------create data structure
type book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Author Author  `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//-------------------declare variables
var booklist = make(map[int]book)
var users = make(map[string]string)
var Router = mux.NewRouter()
var wait time.Duration
var server *http.Server

//Create demo DB
func CreateDB() {

	booklist[1] = book{
		ID:    1,
		Title: "First Book",
		Price: 2243.23,
		Author: Author{
			Firstname: "Kamol",
			Lastname:  "Hasan",
		},
	}

	booklist[2] = book{
		ID:    2,
		Title: "Second Book",
		Price: 23.23,
		Author: Author{
			Firstname: "Masudur",
			Lastname:  "Rahman",
		},
	}

	booklist[3] = book{
		ID:    3,
		Title: "Third Book",
		Price: 243.23,
		Author: Author{
			Firstname: "Rez1",
			Lastname:  "t",
		},
	}

	//user name and password added
	users["admin"] = "admin"
	users["kamol"] = "hasan"

}

//Handle requests
func handleRequest() {

	//init router

	Router.HandleFunc("/books", GetBooks).Methods("GET")
	Router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	Router.HandleFunc("/books", CreateBook).Methods("POST")
	Router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	Router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

}

// delete a certain book and print the rest books
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if value, flag := AuthN(r); !flag {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: " + value))
		return
	}

	params := mux.Vars(r)

	v, _ := strconv.Atoi(params["id"])
	if _, flag := booklist[v]; flag {
		delete(booklist, v)
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Error: Doesn't exist!\n"))
		return
	}

	json.NewEncoder(w).Encode(booklist)
}

// update an existing book with new one
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if value, flag := AuthN(r); !flag {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: " + value))
		return
	}

	params := mux.Vars(r)
	v, _ := strconv.Atoi(params["id"])

	var newbook book
	_ = json.NewDecoder(r.Body).Decode(&newbook)

	if _, flag := booklist[newbook.ID]; !flag {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	booklist[v] = newbook
	json.NewEncoder(w).Encode(booklist)
}

// insert new book info to database
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if value, flag := AuthN(r); !flag {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: " + value))
		return
	}

	var newbook book
	_ = json.NewDecoder(r.Body).Decode(&newbook)

	if _, flag := booklist[newbook.ID]; flag {

		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	booklist[newbook.ID] = newbook
}

// print the info of a certain book
func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if value, flag := AuthN(r); !flag {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: " + value))
		return
	}

	params := mux.Vars(r)
	v, _ := strconv.Atoi(params["id"])
	if value, flag := booklist[v]; flag {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(value)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: 404 Not Found\n"))
	}

}

//print all books listed on database
func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if value, flag := AuthN(r); !flag {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: " + value))
		return

	}

	json.NewEncoder(w).Encode(booklist)

}

//Basic authentication function
func AuthN(r *http.Request) (string, bool) {
	userName, password, flag := r.BasicAuth()

	if flag {

		if value, f := users[userName]; f {

			if value == password {

				return "", true

			} else {

				return "Wrong password!\n", false

			}
		} else {

			return "User doesn't exist!\n", false

		}

	} else {

		return "Header format error!\n", false
	}
}

func CreateSever() {
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	server = &http.Server{
		Addr: ":8000",

		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      Router, 
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

}

func GracefulShutDown() {

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func init() {

	CreateDB()
	handleRequest()

}
