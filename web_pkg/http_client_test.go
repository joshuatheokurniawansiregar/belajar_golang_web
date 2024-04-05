package web_pkg

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestUrlTest(t *testing.T) {
	response, err := http.Get("https://www.python.org/")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// responseio, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(responseio))
	var out *strings.Builder = &strings.Builder{}
	var buff []byte = make([]byte, 1024)
	for {
		n, _ := response.Body.Read(buff)
		if n <= 0 {
			break
		}
		out.Write(buff[:n])
	}
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("SECOND : \n\n\n", out.String())
}
