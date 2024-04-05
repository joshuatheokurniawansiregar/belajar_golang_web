package web_pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name_query := r.URL.Query().Get("name")
	if name_query == "" {
		fmt.Fprint(w, "hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name_query)
	}
}

func MultipleParameter(w http.ResponseWriter, r *http.Request) {
	var firstName string = r.URL.Query().Get("firstname")
	var lastName string = r.URL.Query().Get("lastname")
	if firstName == "" && lastName == "" {
		fmt.Println("")
	} else {
		fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
	}
}
func MultipleParameterValuesQuery(w http.ResponseWriter, r *http.Request) {
	var query url.Values = r.URL.Query()
	var names []string = query["name"]
	fmt.Fprintf(w, "%s", strings.Join(names, " "))

	var name = r.URL.Query().Get("name")
	fmt.Fprintf(w, "%s", name)
}

func TestQueryParameter(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "http://127.0.0.1/hello?name=Joshua", nil)
	responserecorder := httptest.NewRecorder()
	SayHello(responserecorder, request)
	response := responserecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestMultipleQueryParameter(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "http://127.0.0.1/hello?firstname=JoshuaTheo&lastname=Kurniawan", nil)
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	MultipleParameter(responseRecorder, request)
	var response *http.Response = responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestMultipleParameterValues(t *testing.T) {
	var request *http.Request = httptest.NewRequest("GET", "http://127.0.0.1/hello?name=Joshua&name=Theo", nil)
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	MultipleParameterValuesQuery(responseRecorder, request)
	var response *http.Response = responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body: " + string(body) + "\n")
}
