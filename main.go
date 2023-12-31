package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"lessercodes.com/lenslocked/controllers"
	"lessercodes.com/lenslocked/templates"
	"lessercodes.com/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	tmpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tmpl))

	usersController := controllers.Users{}
	usersController.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersController.New)
	r.Post("/users", usersController.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
