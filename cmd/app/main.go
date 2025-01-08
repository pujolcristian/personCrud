package main

import (
	"context"
	"log"
	"personCrud/internal/di"

	"go.uber.org/fx"
)

// @title Person CRUD API
// @version 1.0
// @description API para gestionar personas con operaciones CRUD.
// @host localhost:8080
// @BasePath /
func main() {
	app := fx.New(
		di.ProvideModules,
		di.ProvideMiddlewares,
		di.InvokeModules,
	)

	ctx := context.Background()

	// Iniciar la aplicación
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Error al iniciar la aplicación: %v", err)
	}

	<-app.Done()

	// Detener la aplicación
	if err := app.Stop(ctx); err != nil {
		log.Fatalf("Error al detener la aplicación: %v", err)
	}
}
