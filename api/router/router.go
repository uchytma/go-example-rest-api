package router

import (
	"log"
	"net/http"
	"time"

	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/uchytma/go-example-rest-api/api/endpoints/math"
	_ "github.com/uchytma/go-example-rest-api/docs"
)

func Initialize() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), //forces Content-type
		middleware.RedirectSlashes,
		middleware.Recoverer,            //middleware to recover from panics
		middleware.Heartbeat("/health"), //for heartbeat process such as Kubernetes liveprobeness
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	)

	fs := getHandlerWithDoc()

	//Sets context for all requests
	router.Use(middleware.Timeout(30 * time.Second))
	if fs != nil {
		router.Handle("/swagger/json/swagger.json", http.StripPrefix("/swagger/json/", fs))
	}
	router.Route("/", func(r chi.Router) {
		r.Mount("/math", math.Routes())
	})

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/json/swagger.json"), //The url pointing to API definition
	))

	return router
}

func getHandlerWithDoc() http.Handler {
	//detect first exists directory. There should be swagger.json.
	docSearchLocations := []string{"docs", "doc"}

	for _, location := range docSearchLocations {
		if _, err := os.Stat(location); !os.IsNotExist(err) {
			return http.FileServer(http.Dir(location))
		}
	}
	return nil
}

func ServeRouter() {
	r := Initialize()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error serving router")
	}
}
