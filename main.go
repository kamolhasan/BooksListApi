package main

import (
	"github.com/kamolhasan/BookListApi/api"
	"log"
	"net/http"

)

func main()  {
	log.Fatal(http.ListenAndServe(":8000",api.Router))

}