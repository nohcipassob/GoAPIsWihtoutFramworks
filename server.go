package main

import (
	"log"
	"net/http"
)

// type
type APIServer struct {
	addr string
}

type Middleware func(http.Handler) http.HandlerFunc

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {

	router := http.NewServeMux()
	router.HandleFunc("POST /users/{userId}",  func(w http.ResponseWriter, r *http.Request){

		userId := r.PathValue("userId")
		w.Write([]byte("User ID: " + userId))
	})

	router.HandleFunc("/users/{userId}",  func(w http.ResponseWriter, r *http.Request){

		w.Write([]byte("CATCH ALL METHOD"))
	})

	// Apply middleware stacks
	MiddlewareChain := MiddlewareStacks(AuthMiddleware, RequestLoggerMiddleware)

	server := http.Server{
		Addr: s.addr,
		Handler: MiddlewareChain(router),
	}

	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()

}

// Request loggin middleware
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s, path: %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
} 

// Authen middleware
func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check token
	token := r.Header.Get("Authorization")
	if token != "Bearer token" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	next.ServeHTTP(w, r)
	}
}

func MiddlewareStacks(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}

		return next.ServeHTTP
	}
}