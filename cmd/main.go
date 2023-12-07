package main

import (
	"checkout/api/config"
	"net/http"
)

func main() {
	var router = config.Create_engine()
	http.ListenAndServe(":9090", router)
}
