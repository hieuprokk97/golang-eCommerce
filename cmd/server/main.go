package main

import (
	"github.com/hieuprokk97/golang-eCommerce.git/internal/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":8002")
}
