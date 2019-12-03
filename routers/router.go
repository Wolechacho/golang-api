package routers

import (
	"fmt"
	"net/http"

	"first-api-golang/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	pc := &controllers.ProductController{}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/product", pc.Save).Methods("POST")
	// router.HandleFunc("/api/product", handler).Methods("GET")
	// router.HandleFunc("/api/product/{id}", handler).Methods("GET")
	// router.HandleFunc("/api/product/{id}", handler).Methods("DELETE")
	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
