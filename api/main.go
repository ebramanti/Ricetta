package main

import (
	"flag"
	"fmt"
	a "github.com/jadengore/Ricetta/api/api"
	"github.com/jadengore/Ricetta/api/routes"
	"github.com/jadengore/goconfig"
	"log"
	"net/http"
)

// Flag for testing.
var testing = flag.Bool("testing", false, "For API testing")

func main() {
	config, _ := goconfig.ReadConfigFile("../api-config.cfg")
	port, _ := config.GetString("default", "server-port")
	uri, _ := config.GetString("local", "db-url")

	api := a.NewApi(uri, config)
	handler, err := routes.MakeHandler(*api, *testing)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("docs"))))

	http.Handle("/v1/", http.StripPrefix("/v1", &handler))

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
