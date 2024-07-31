package main

import (
	"teacher2/model"
	"teacher2/routers"
)

func main() {

	model.InitDb()
	routers.InitRouter()

}
