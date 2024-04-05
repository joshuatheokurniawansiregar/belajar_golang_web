package web_pkg

import (
	"fmt"
	"net/http"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	var cookie *http.Cookie = new(http.Cookie)
	cookie.Name = "X-Joshua-Name"
	cookie.Value = request.URL.Query().Get("name")
	http.SetCookie(writer, cookie)
	cookie.Path = "/"
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Joshua-Name")
	if err != nil {
		panic(cookie)
	} else {
		var name string = cookie.Name
		fmt.Fprintf(writer, "%s", name)
	}

}
