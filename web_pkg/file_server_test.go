package web_pkg

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed all:resources
var resources_new embed.FS

func TestFileServerWithoutEmbed(t *testing.T) {
	directory := http.Dir("../resources")
	var fileServer http.Handler = http.FileServer(directory)
	fmt.Printf("%s", fileServer)
	var serveMux *http.ServeMux = http.NewServeMux()
	serveMux.Handle("/static/", fileServer)

	var server *http.Server = &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestFileServerUseEmbed(t *testing.T) {
	directory, err_ := fs.Sub(resources_new, "resources")
	if err_ != nil {
		panic(err_)
	}
	var fileServer http.Handler = http.FileServer(http.FS(directory))
	var serveMux *http.ServeMux = http.NewServeMux()
	serveMux.Handle("/static/", http.StripPrefix("/static", fileServer))

	var server *http.Server = &http.Server{
		Addr:    ":8081",
		Handler: serveMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
