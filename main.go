package main

import (
	db "assignment-2_GiandyRM/database"
	"assignment-2_GiandyRM/routers"
)

func main() {
	db.Init()
	routers.ServerOn().Run(":8080")
}
