package main

import (
	"fmt"
	"log"
	"net/http"

	"HENNGE/gip-interns/CodeSamples/go-hello-web/dice"
	"HENNGE/gip-interns/CodeSamples/go-hello-web/internal/webservice"
)

const listenPort = 8080

func main() {
	diceSDK := dice.New()
	service := webservice.New(diceSDK)

	http.HandleFunc("/ping", service.Ping)
	http.HandleFunc("/roll", service.Roll)

	log.Printf("Listening on %d\n", listenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil))
}
