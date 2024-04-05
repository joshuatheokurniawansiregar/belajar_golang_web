package main

import (
	"belajar_golang_web/web_pkg"
	"log"
	"net/http"
)

func RunTestCookie() {
	var serveMux *http.ServeMux = http.NewServeMux()
	serveMux.HandleFunc("/set-cookie", web_pkg.SetCookie)
	serveMux.HandleFunc("/get-cookie", web_pkg.GetCookie)
	var webServer http.Server = http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	log.Fatal(webServer.ListenAndServe())
}

func main() {
	RunTestCookie()
}
