package main

import (
	"log"
	"net/http"
)

func main() {
	// se configura para leer archivos html desde la carpeta 'static'
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// iniciar servidor
	log.Println("escuchando en puerto 8080:")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}