package web_pkg

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"testing"
)

func SimplessHtml(w http.ResponseWriter, r *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "SIMPLE", "hello world")
}

func TestTemplate(t *testing.T) {
	var serveMux *http.ServeMux = http.NewServeMux()
	serveMux.HandleFunc("/", SimplessHtml)
	var server *http.Server = &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	log.Fatal(server.ListenAndServe())
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.html"))
	t.ExecuteTemplate(w, "simple.html", "TestTemplate")
}

func TestSimpleHTMLFileUsingDir(t *testing.T) {
	var dir = http.Dir("./resources")
	var serveMux *http.ServeMux = http.NewServeMux()
	var fileServer http.Handler = http.FileServer(dir)
	serveMux.HandleFunc("/", SimpleHTMLFile)
	serveMux.Handle("/static/", http.StripPrefix("/static", fileServer))
	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	log.Fatal(server.ListenAndServe())
}

func TestSimpleHTMLFile(t *testing.T) {
	// directory, err_ := fs.Sub(resources_new, "resources")
	// if err_ != nil {
	// 	panic(err_)
	// }
	var serveMux *http.ServeMux = http.NewServeMux()
	var fileServer http.Handler = http.FileServer(http.FS(resources_new))
	serveMux.HandleFunc("/", SimpleHTMLFile)
	serveMux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	var server http.Server = http.Server{
		Addr:    ":8081",
		Handler: serveMux,
	}
	log.Fatal(server.ListenAndServe())
}

//go:embed dist/simple.html
var simple_html embed.FS

var myTemplates = template.Must(template.ParseFS(simple_html, "dist/simple.html"))

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.html", map[string]interface{}{
		"Name":    r.URL.Query().Get("name"),
		"Address": "Test Address",
	})
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}
	log.Fatal(server.ListenAndServe())
}

//go:embed dist/*.html
var templates embed.FS

var templatePages *template.Template = template.Must(template.ParseFS(templates, "dist/*.html"))

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	templatePages.ExecuteTemplate(w, "blog.html", "")
}

func TestTemplateEmbed(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
