package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Router to handle requests in the backend*/
var router *mux.Router

/*StartRouter setup an HTTP server in the configured port and api base path*/
func StartRouter() *mux.Router {
	router = mux.NewRouter()

	router.Use(loggingMiddleware)

	handler:= setupCors(router)
	router.Handle("/", handler)

	return router
}

func setupCors(router *mux.Router) http.Handler {
	corsOptions := cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowCredentials: true,
	}
	cors:= cors.New(corsOptions)

	handler := cors.Handler(router)

  	return handler
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}