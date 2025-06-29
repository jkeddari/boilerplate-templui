package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/jkeddari/boilerplate-templui/ui/pages"
	"github.com/joho/godotenv"
)

const defaultAppPort = "8080"

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	port   string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = defaultAppPort
	}
}

func main() {
	mux := http.NewServeMux()
	setupAssetsRoutes(mux)

	mux.Handle("GET /", templ.Handler(pages.Index("boilerplate-templui")))

	logger.Info("Server is running", "port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func setupAssetsRoutes(mux *http.ServeMux) {
	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var fs http.Handler
		fs = http.FileServer(http.Dir("./assets"))
		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
