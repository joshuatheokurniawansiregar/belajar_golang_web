package web_pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello worldssss")
}

func TestHTTP(t *testing.T) {
	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/hello", nil)
	responserecorder := httptest.NewRecorder()
	handler(responserecorder, request)
	response := responserecorder.Result()
	bodyresponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyresponse))
	fmt.Println(response.Status)
	fmt.Println(response.Status)
}
