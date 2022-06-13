package storage

import (
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
)

func GetInventoryById(inventoryId string) entities.Inventory {
	var inventory entities.Inventory
	database.Instance.First(&inventory, inventoryId)
	return inventory
}

func GetInventories(showDeleted bool) []entities.Inventory {
	var inventories []entities.Inventory
	if showDeleted {
		database.Instance.Find(&inventories)
	} else {
		database.Instance.Where("deleted = ?", false).Find(&inventories)
	}
	return inventories
}

func CreateInventory(inventory entities.Inventory) {
	database.Instance.Create(&inventory)
}

func SaveInventory(inventory entities.Inventory) {
	database.Instance.Save(&inventory)
}
