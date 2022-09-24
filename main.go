package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Printf("parsing templates: %v", err)
		http.Error(w, "There was an error parsing the templates", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("excuting templates: %v", err)
		http.Error(w, "There was an error excuting the templates", http.StatusInternalServerError)
		return
	}

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"gideon10@gmail.com\">gideon@gmail</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li><strong>Is there a free version?</strong> Yes! We offer a free trial for 30 days on any paid plans</li>
		<li><Strong>What are your support hours?</strong> We have support staff answering emails 24/7, though response time may be a bit slower on weekens.</li>
		<li><strong>How do I contact support?</strong></li> Email us - <a href="gideon10@gmail.com" >support@gideon</a></li>
	</ul>
	`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homehandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
