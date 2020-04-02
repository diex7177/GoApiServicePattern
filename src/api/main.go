package main

import (
	router "GoApiServicePattern/src/api/router"
)

func main() {
	r := router.InitRouter()
	r.Run()
}
