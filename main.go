package main

import (
	"TikTok/database"
	"TikTok/router"
)

func main() {
	database.Init()
	router.SetupRouter()
	router.Router.Run()
}
