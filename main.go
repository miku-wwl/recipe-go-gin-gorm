package main

import (
	"recipe/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
}
