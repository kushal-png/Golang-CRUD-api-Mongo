package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kushal-png/mongoapi/router"
)

func main() {
	fmt.Println("MOngoDb API")
	fmt.Println("Server is starting")
    r:= router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on 4000")
}
