package app

import (
	"net/http"

	"github.com/adhiardiansyah/goshop/app/controllers"
	"github.com/adhiardiansyah/goshop/middleware"
	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/login", controllers.Login).Methods("GET")
	server.Router.HandleFunc("/processlogin", controllers.ProcessLogin).Methods("POST")
	server.Router.HandleFunc("/welcome", middleware.LoginMiddleware(controllers.Welcome)).Methods("GET")
	server.Router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	staticFileDirectory := http.Dir("./assets")
	staticFileHandler := http.StripPrefix("/public", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public").Handler(staticFileHandler).Methods("GET")
}
