package main

import (
	"log"
	"net/http"

)

func main() {
	// se configura para leer archivos html desde la carpeta 'static'
	fs := http.FileServer(http.Dir("static"))

	// se usa el metodo StripPrefix para servir index.html como default en /
	http.Handle("/", http.StripPrefix("/", fs))
	
	// iniciar servidor
	log.Println("escuchando en puerto 8080:")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}