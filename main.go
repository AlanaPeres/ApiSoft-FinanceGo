package main

import (
	"ApiSoft-Finance/src/config"
	"ApiSoft-Finance/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Api Soft-Finance")
	config.Carregar()
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
