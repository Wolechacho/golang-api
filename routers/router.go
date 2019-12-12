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
	sc := &controllers.CategoryController{}
	oc := &controllers.OrderController{}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	//api routes for product
	router.HandleFunc("/api/product", pc.Save).Methods("POST")
	router.HandleFunc("/api/product", pc.List).Methods("GET")
	router.HandleFunc("/api/product/{id}", pc.Get).Methods("GET")
	router.HandleFunc("/api/product/{id}", pc.Delete).Methods("DELETE")

	//api routes for category
	router.HandleFunc("/api/category", sc.Save).Methods("POST")

	//api routes for order
	router.HandleFunc("/api/order", oc.Save).Methods("POST")

	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
