package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RegisterInventoryRoutes(router *mux.Router) {
	router.HandleFunc("/api/inventory", controllers.GetInventories).Methods("GET")
	router.HandleFunc("/api/inventory/{id}", controllers.GetInventoryById).Methods("GET")
	router.HandleFunc("/api/inventory", controllers.CreateInventory).Methods("POST")
	router.HandleFunc("/api/inventory/{id}", controllers.UpdateInventory).Methods("PUT")
	router.HandleFunc("/api/inventory/{id}", controllers.DeleteInventory).Methods("DELETE")
	router.HandleFunc("/api/inventory/{id}/undelete", controllers.UndeleteInventory).Methods("POST")
}

func main() {
	LoadAppConfig()

	database.Connect(os.Getenv("CONNECTION_STRING"))
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterInventoryRoutes(router)

	log.Printf("Starting server on port %v", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
