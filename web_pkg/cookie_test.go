package web_pkg

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func SetCookieTest(writer http.ResponseWriter, request *http.Request) {
	var cookie *http.Cookie = new(http.Cookie)
	cookie.Name = "X-Joshua-Name"
	cookie.Value = request.URL.Query().Get("name")
	http.SetCookie(writer, cookie)
	cookie.Path = "/"
}

func GetCookieTest(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Joshua-Name")
	if err != nil {
		panic(cookie)
	} else {
		var name string = cookie.Name
		var value string = cookie.Value
		fmt.Fprintf(writer, "name: %s. Value: %s", name, value)
	}

}

func TestCookie(t *testing.T) {
	var serveMux *http.ServeMux = http.NewServeMux()
	serveMux.HandleFunc("/set-cookie", SetCookieTest)
	serveMux.HandleFunc("/get-cookie", GetCookieTest)
	var webServer http.Server = http.Server{
		Addr:    ":8083",
		Handler: serveMux,
	}
	log.Fatal(webServer.ListenAndServe())
}
