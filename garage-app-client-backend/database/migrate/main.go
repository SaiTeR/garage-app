package main

import (
	"fmt"
	"garage-app-client-backend/database"
)

func main() {
	database.InitMySQL()
	database.MigrateUsers()
	fmt.Print("Migration completed!")
}
