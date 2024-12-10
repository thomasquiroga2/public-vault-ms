package main

import (
	"log"
	"net/http"

	"public-vault-ms/config"
	"public-vault-ms/controllers"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar el controlador
	http.HandleFunc("/tokenize", controllers.TokenizeCardHandler)
	http.HandleFunc("/detokenize", controllers.DetokenizeCardHandler)

	// Iniciar el servidor
	log.Printf("Servidor corriendo en el puerto %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
