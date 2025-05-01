package setup

import (
	utils "web_app_fiber/helpers"
)

func InitSetup() {
	CreateDb()
	CreateTables()
	utils.InitializeConnectionPool("main")
}
