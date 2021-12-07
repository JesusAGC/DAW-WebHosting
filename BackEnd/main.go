package main

import (
	api "BackendOrdinario/API"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := api.MyRoutes()
	log.Println("Servidor en funcionamiento...")
	fmt.Println(time.Now())
	log.Fatal(http.ListenAndServe(":9000", router))

}
