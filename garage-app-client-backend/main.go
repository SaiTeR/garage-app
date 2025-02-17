package main

import "garage-app-client-backend/database"

func main() {
	database.InitMySQL()
	database.InitRedis()

	// Получаем настроенный роутер из SetupRouter
	r := SetupRouter()

	// Запускаем сервер на порту 8081
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}

// docker compose -p garage-app up -d --build
