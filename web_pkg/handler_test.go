package web_pkg

import (
	"fmt"
	"net/http"
	"testing"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func TestServeMux(t *testing.T) {
	var serveMux *http.ServeMux = http.NewServeMux()
	var handler http.HandlerFunc = func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "hello world")
	}
	serveMux.HandleFunc("/aye", handler)
	serveMux.HandleFunc("/", home)
	serveMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	serveMux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "images page")
	})
	serveMux.HandleFunc("/images/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "thumbnail page")
	})
	var server *http.Server = &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic("ada error bang: " + err.Error())
	}
}
