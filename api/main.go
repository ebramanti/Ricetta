package main

import (
	routes "./routes"
	"fmt"
	"github.com/jadengore/goconfig"
	"log"
	"net/http"
)

func main() {
	config, err := goconfig.ReadConfigFile("config.cfg")
	port, err := config.GetString("default", "server-port")

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
