FROM golang:latest
RUN go get -u github.com/spf13/cobra/cobra
RUN go get -u github.com/gorilla/mux
COPY . /go/src/github.com/kamolhasan/BookListApi
WORKDIR /go/src/github.com/kamolhasan/BookListApi
RUN go build
ENTRYPOINT ["/go/src/github.com/kamolhasan/BookListApi/BookListApi"]
EXPOSE 8000

