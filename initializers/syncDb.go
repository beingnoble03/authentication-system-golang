package initializers

import models "github.com/beingnoble03/octern-main/models"

func SyncDb() {
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Organization{})
	Db.AutoMigrate(&models.OrganizationUser{})
}
