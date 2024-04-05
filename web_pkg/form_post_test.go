package web_pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	var err error = request.ParseForm()
	if err != nil {
		panic(err)
	}
	var firstName string = request.PostForm.Get("first_name")
	var lastName string = request.PostForm.Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	var stringReader io.Reader = strings.NewReader("first_name=Joshua Theo&last_name=Kurniawan Siregar")
	var request *http.Request = httptest.NewRequest("POST", "http://127.0.0.1:8080/hello", stringReader)
	request.Header.Add("content-Type", "application/x-www-form-urlencoded")
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	FormPost(responseRecorder, request)
	var response *http.Response = responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
