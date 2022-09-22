package main

import (
	"fmt"
	"net/http"
	"storage-api/api"
)

func main() {

	fmt.Println("running....")
	api.ServerHTTP()

	const port = "8282"

	http.ListenAndServe(":"+port, nil)
	fmt.Printf("running on port %s", port)
}
