package api

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Router to handle requests in the backend*/
var Router *mux.Router

/*StartRouter setup an HTTP server in the configured port and api base path*/
func StartRouter() *mux.Router {
	Router = mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler:= setupCors(Router)

	http.ListenAndServe(":"+PORT, handler)

	return Router
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