package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// delete a certain book and print the rest books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if value,flag:=AuthN(r);!flag{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: "+value))
		return
	}


	params := mux.Vars(r)

	v, _ := strconv.Atoi(params["id"])
	if _, flag := booklist[v]; flag {
		delete(booklist, v)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Doesn't exist!\n"))
		return
	}

	json.NewEncoder(w).Encode(booklist)
}

// update an existing book with new one
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if value,flag:=AuthN(r);!flag{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: "+value))
		return
	}

	params := mux.Vars(r)
	v, _ := strconv.Atoi(params["id"])

	var newbook book
	_ = json.NewDecoder(r.Body).Decode(&newbook)
	booklist[v] = newbook

	json.NewEncoder(w).Encode(booklist)
}

// insert new book info to database
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if value,flag:=AuthN(r);!flag{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: "+value))
		return
	}

	w.WriteHeader(http.StatusCreated)
	var newbook book
	_ = json.NewDecoder(r.Body).Decode(&newbook)
	booklist[newbook.ID] = newbook
}

// print the info of a certain book
func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if value,flag:=AuthN(r);!flag{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: "+value))
		return
	}


	params := mux.Vars(r)
	v, _ := strconv.Atoi(params["id"])
	if value, flag := booklist[v]; flag {
		json.NewEncoder(w).Encode(value)
		return
	}else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: 400 Bad Request\n"))
	}

}

//print all books listed on database
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if value,flag:=AuthN(r);!flag{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: "+value))
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

func main() {

	CreateDB()
	handleRequest()

}
