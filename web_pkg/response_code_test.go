package web_pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HandlerResponseCode(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	var firstName string = request.PostForm.Get("first_name")
	var lastName string = request.PostForm.Get("last_name")
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Not Allowed Method %s", request.Method)
	} else {
		if firstName == "" && lastName == "" {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "%s %s", firstName, lastName)
		} else {
			fmt.Fprintf(writer, "%s %s", firstName, lastName)
		}
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	var inputReader io.Reader = strings.NewReader("first_name=Joshua Theo&last_name=Kurniawan Siregar")
	var request *http.Request = httptest.NewRequest(http.MethodPost, "http://127.0.0.1/hello", inputReader)
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	request.Header.Add("content-Type", "application/x-www-form-urlencoded")
	HandlerResponseCode(responseRecorder, request)
	var response = responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(body))
	fmt.Printf("%d\n", response.StatusCode)
}
