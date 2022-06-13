package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/entities"
	"golang-crud-rest-api/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Main(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		"Server is running! Proceed to https://github.com/alyoanton9/inventory-tracking-crud for information on API")
}

func GetInventories(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	showDeletedParam := r.URL.Query().Get("show_deleted")
	showDeleted := false

	showDeleted, parseError := strconv.ParseBool(showDeletedParam)
	if showDeletedParam != "" && parseError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inventories := storage.GetInventories(showDeleted)

	json.NewEncoder(w).Encode(inventories)
}

func GetInventoryById(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	inventoryId := mux.Vars(r)["id"]
	if !inventoryExist(inventoryId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	inventory := storage.GetInventoryById(inventoryId)

	if inventory.Deleted {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(inventory)
}

func CreateInventory(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	var inventory entities.Inventory
	json.NewDecoder(r.Body).Decode(&inventory)

	storage.CreateInventory(inventory)
	json.NewEncoder(w).Encode(inventory)
}

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	inventoryId := mux.Vars(r)["id"]
	if !inventoryExist(inventoryId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	inventory := storage.GetInventoryById(inventoryId)
	json.NewDecoder(r.Body).Decode(&inventory)

	storage.SaveInventory(inventory)
	json.NewEncoder(w).Encode(inventory)
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	inventoryId := mux.Vars(r)["id"]
	if !inventoryExist(inventoryId){
		w.WriteHeader(http.StatusNotFound)
		return
	}

	inventory := storage.GetInventoryById(inventoryId)
	if inventory.Deleted {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	inventory.Deleted = true
	inventory.Comment = r.URL.Query().Get("comment")

	storage.SaveInventory(inventory)
	json.NewEncoder(w).Encode(inventory)
}

func UndeleteInventory(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	inventoryId := mux.Vars(r)["id"]
	if !inventoryExist(inventoryId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	inventory := storage.GetInventoryById(inventoryId)
	if !inventory.Deleted {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	inventory.Deleted = false
	inventory.Comment = ""

	storage.SaveInventory(inventory)
	json.NewEncoder(w).Encode(inventory)
}

func inventoryExist(inventoryId string) bool {
	inventory := storage.GetInventoryById(inventoryId)
	if inventory.Id == 0 {
		return false
	}
	return true
}

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
