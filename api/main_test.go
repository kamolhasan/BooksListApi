package api

import (
	"encoding/base64"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Request struct {
	Method             string
	URL                string
	Body               io.Reader
	ExpectedStatusCode int

}

func TestGetBooks(t *testing.T) {
	requests:=make([]Request,3)

	requests[0]=Request{
		"GET",
		"http://localhost:8000/books" ,
		nil,
		200,

	}
	requests[1]=Request{
		"GET",
		"http://localhost:8000/books" ,
		nil,
		201,

	}
	requests[2]=Request{
		"GET",
		"http://localhost:8000" ,
		nil,
		200,

	}

	processRequest(t,requests)
}


func processRequest(t *testing.T, reqs []Request) {
	for _, req := range reqs {
		r, _ := http.NewRequest(req.Method, req.URL, req.Body)
		r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("kamol:hasan")))
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, r)
		if w.Code != req.ExpectedStatusCode{
			t.Error("\nExpected Status Code\t= "+cast.ToString(req.ExpectedStatusCode)+"\nFound Status Code\t\t= "+cast.ToString(w.Code)+"\n")

		}


	}
}
