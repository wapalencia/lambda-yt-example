package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, req Request) (Response, error) {
	// Procesar la solicitud y generar una respuesta
	message := fmt.Sprintf("Hola, %s. ¡Bienvenido a AWS Lambda en Go!", req.Name)
	response := Response{
		Message: message,
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
