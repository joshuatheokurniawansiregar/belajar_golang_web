package web_pkg

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type testhandler struct {
}

func (t *testhandler) anotherhandler(w http.ResponseWriter, r *http.Request) {

}

func TestServer(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "hello world")
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// var objtest *testhandler = &testhandler{}

	// another_server := http.Server{
	// 	Addr:    ":8081",
	// 	Handler: objtest.anotherhandler,
	// }
	log.Fatal(server.ListenAndServe())
}
