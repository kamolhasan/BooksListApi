package api

import (
	"encoding/base64"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Request struct {
	Method             string
	URL                string
	Body               io.Reader
	ExpectedStatusCode int
}

func TestCreateBook(t *testing.T) {
	requests := make([]Request, 2)
	requests[0] = Request{
		"POST",
		"http://localhost:8000/books",
		strings.NewReader(`{"id":1,"title":"First Book","price":2243.23,"author":{"firstname":"Kamol","lastname":"Hasan"}}`),
		405,
	}

	requests[1] = Request{
		"POST",
		"http://localhost:8000/books",
		strings.NewReader(`{"id":4,"title":"First Book","price":2243.23,"author":{"firstname":"Kamol","lastname":"Hasan"}}`),
		201,
	}

	processRequest(t, requests)
}

func TestGetBook(t *testing.T) {
	requests := make([]Request, 3)

	requests[0] = Request{
		"GET",
		"http://localhost:8000/books/3",
		nil,
		200,
	}
	requests[1] = Request{
		"GET",
		"http://localhost:8000/books/5",
		nil,
		404,
	}
	requests[2] = Request{
		"GET",
		"http://localhost:8000/book/2",
		nil,
		404,
	}

	processRequest(t, requests)
}

func TestGetBooks(t *testing.T) {
	requests := make([]Request, 3)

	requests[0] = Request{
		"GET",
		"http://localhost:8000/books",
		nil,
		200,
	}
	requests[1] = Request{
		"GET",
		"http://localhost:8000/book",
		nil,
		404,
	}
	requests[2] = Request{
		"GET",
		"http://localhost:8000/books",
		nil,
		200,
	}

	processRequest(t, requests)
}

func TestUpdateBook(t *testing.T) {
	requests := make([]Request, 2)
	requests[0] = Request{
		"PUT",
		"http://localhost:8000/books/2",
		strings.NewReader(`{"id":2,"title":"First Book","price":2243.23,"author":{"firstname":"Kamol","lastname":"Hasan"}}`),
		200,
	}

	requests[1] = Request{
		"PUT",
		"http://localhost:8000/books/5",
		strings.NewReader(`{"id":5,"title":"First Book","price":2243.23,"author":{"firstname":"Kamol","lastname":"Hasan"}}`),
		204,
	}
	processRequest(t,requests)
}

func TestDeleteBook(t *testing.T) {
	requests := make([]Request, 2)
	requests[0] = Request{
		"DELETE",
		"http://localhost:8000/books/1",
		nil,
		200,
	}

	requests[1] = Request{
		"DELETE",
		"http://localhost:8000/books/1",
		nil,
		204,
	}

	processRequest(t, requests)
}

func processRequest(t *testing.T, reqs []Request) {
	for _, req := range reqs {
		r, _ := http.NewRequest(req.Method, req.URL, req.Body)
		r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("kamol:hasan")))
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, r)
		if w.Code != req.ExpectedStatusCode {
			t.Error("\nExpected Status Code\t= " + cast.ToString(req.ExpectedStatusCode) + "\nFound Status Code\t\t= " + cast.ToString(w.Code) + "\n")

		}

	}
}
