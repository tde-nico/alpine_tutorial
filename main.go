package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"alpine-tutorial/internal/middleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// Parse templates
	tmpl := template.Must(template.New("").ParseGlob("templates/*.html"))

	r := http.NewServeMux()

	// Handle Static Files
	static := http.FileServer(http.Dir("./static"))
	r.Handle("GET /static/", http.StripPrefix("/static", static))

	// Handle routes
	r.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "menu.html", nil)
	})

	r.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "todo.html", nil)
	})

	r.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "menu.html", nil)
	})

	r.HandleFunc("/collapse", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "collapse.html", nil)
	})

	srv := &http.Server{
		Handler: middleware.Logging(logger)(r),
		Addr:    ":8081",
	}

	logger.Info("server listening", slog.String("addr", srv.Addr))

	srv.ListenAndServe()
}
