package routers

import (
	"fmt"
	"net/http"

	"first-api-golang/controllers"

	"github.com/gorilla/mux"
)

//RegisterRoutes - Configuration for all incoming routes
func RegisterRoutes() *mux.Router {
	pc := &controllers.ProductController{}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/product", pc.Save).Methods("POST")
	router.HandleFunc("/api/product", pc.List).Methods("GET")
	router.HandleFunc("/api/product/{id}", pc.Get).Methods("GET")
	router.HandleFunc("/api/product/{id}", pc.Delete).Methods("DELETE")
	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
