package initializers

import "github.com/OrdinarilySam/feddit/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
