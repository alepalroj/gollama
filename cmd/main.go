package main

import (
	"fmt"
	"log"
	"github.com/alepalroj/gollama"
)

func main() {
	config, err := gollama.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error en carga de configuración: %v", err)
	}

	client := gollama.OllamaClient(config)

	req := gollama.Request{
		Model:  config.Model,
		Prompt: "Escribe un poema corto",
		Stream: false,
	}

	resp, err := client.Generate(req)
	if err != nil {
		log.Fatalf("Error al generar respuesta: %v", err)
	}

	fmt.Printf("Respuesta: %s\n", resp.Response)
}