package main

import (
	"fmt"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awsome site!</h1>")
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

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homehandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
