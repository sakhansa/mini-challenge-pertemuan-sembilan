package main

import (
	"mini-challenge-pertemuan-sembilan/database"
	"mini-challenge-pertemuan-sembilan/router"
)

var (
	PORT = ":9090"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
