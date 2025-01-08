package di

import (
	"personCrud/internal/adapters/db"
	"personCrud/internal/api"
	"personCrud/internal/api/controllers"
	"personCrud/internal/api/middlewares"
	"personCrud/internal/api/routes"
	"personCrud/internal/domain/usecases"

	"go.uber.org/fx"
)

// ProvideModules TIP <p> Agrupa todas las dependencias del proyecto </p>
var ProvideModules = fx.Options(
	// Proveer dependencias
	fx.Provide(db.ConnectToDB),
	fx.Provide(db.NewGormPersonRepository),
	fx.Provide(usecases.NewPersonUseCase),
	fx.Provide(controllers.NewPersonController),
	fx.Provide(api.NewRouter),
)

// InvokeModules TIP <p> Agrupa todas las funciones a invocar </p>
var InvokeModules = fx.Options(
	fx.Invoke(routes.RegisterPersonRoutes),
	fx.Invoke(api.StartServer),
)

var ProvideMiddlewares = fx.Options(
	fx.Provide(
		fx.Annotate(middlewares.AuthMiddleware, fx.ResultTags(`name:"authMiddleware"`)),
		fx.Annotate(middlewares.CORSMiddleware, fx.ResultTags(`name:"corsMiddleware"`)),
		fx.Annotate(middlewares.LoggingMiddleware, fx.ResultTags(`name:"loggingMiddleware"`)),
	),
)
