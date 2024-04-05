package web_pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	var contentType string = r.Header.Get("content-type")
	fmt.Println(contentType)
}

func ResponseHeader(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Add("X-Powered-By", "Joshua Theo")
}

func TestRequestHeader(t *testing.T) {
	var request *http.Request = httptest.NewRequest("GET", "http://127.0.0.1/hello", nil)
	request.Header.Add("content-type", "application/json")
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	RequestHeader(responseRecorder, request)
	var response = responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestReponseHeader(t *testing.T) {
	var request *http.Request = httptest.NewRequest("GET", "http://127.0.0.1/hello", nil)
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	ResponseHeader(responseRecorder, request)
	var responseRecorderHeader = responseRecorder.Header().Get("X-Powered-bY")
	var response *http.Response = responseRecorder.Result()
	var header = response.Header.Get("X-Powered-by")
	fmt.Println(header)
	fmt.Println(responseRecorderHeader)
}
