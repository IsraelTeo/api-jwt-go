package route

import (
	"github.com/IsraelTeo/api-jwt-go/handler"
	"github.com/IsraelTeo/api-jwt-go/middelware"
	"github.com/gorilla/mux"
)

var (
	loginPath = "/login"

	userBasicPath = "/user"
	userIDPath    = "/user/{id}"
	usersPath     = "/users"

	roleBasicPath = "/role"
	roleIDPath    = "/role/{id}"
	rolesPath     = "/roles"
)

func InitRoute() *mux.Router {

	routes := mux.NewRouter()

	api := routes.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc(loginPath, middelware.SetMiddelwareAuthentication(middelware.Log(handler.GetUserById))).Methods("POST")

	api.HandleFunc(userIDPath, middelware.Log(handler.GetUserById)).Methods("GET")
	api.HandleFunc(usersPath, middelware.Log(handler.GetAllUsers)).Methods("GET")
	api.HandleFunc(userBasicPath, middelware.Log(handler.SaveUser)).Methods("POST")
	api.HandleFunc(userIDPath, middelware.Log(handler.UpdateUser)).Methods("PUT")
	api.HandleFunc(userIDPath, middelware.Log(handler.DeleteUser)).Methods("DELETE")

	api.HandleFunc(roleIDPath, middelware.Log(handler.GetRoleById)).Methods("GET")
	api.HandleFunc(rolesPath, middelware.Log(handler.GetAllRoles)).Methods("GET")
	api.HandleFunc(roleBasicPath, middelware.Log(handler.SaveRole)).Methods("POST")
	api.HandleFunc(roleIDPath, middelware.Log(handler.UpdateRole)).Methods("PUT")
	api.HandleFunc(roleIDPath, middelware.Log(handler.DeleteRole)).Methods("DELETE")

	return routes
}
