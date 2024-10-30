package route

import (
	"fmt"

	"github.com/gorilla/mux"
)

func InitRoute() *mux.Router {

	routes := mux.NewRouter()
	api := routes.PathPrefix("/api/v1").Subrouter()

	fmt.Println(api)
	return routes
}
