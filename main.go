package main

import (
	"fmt"
	"net/http"

	"github.com/Gideon-isa/lenslocked/controllers"
	"github.com/Gideon-isa/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Parse the template
	r := chi.NewRouter()
	//
	tpl := views.Must(views.Parse("templates/home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	//
	tpl = views.Must(views.Parse("templates/contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	//
	tpl = views.Must(views.Parse("templates/faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
