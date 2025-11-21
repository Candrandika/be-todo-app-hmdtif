package bootstrap

import (
	"log"

	"github.com/Candrandika/be-todo-app-hmdtif/config"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/repository"
	"github.com/Candrandika/be-todo-app-hmdtif/domain/usecase"
	"github.com/Candrandika/be-todo-app-hmdtif/infra/db"
	"github.com/Candrandika/be-todo-app-hmdtif/interface/handler"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Application struct {
	Env      *config.Env
	DB       *gorm.DB
	FiberApp *fiber.App
}

func App() *Application {
	app := &Application{}
	app.Env = config.NewEnv()
	app.DB = db.NewPostgresDatabase(app.Env)
	db.Migrate(app.DB)

	app.FiberApp = fiber.New()

	taskRepo := repository.NewTaskRepository(app.DB)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase, validator.New())

	tasks := app.FiberApp.Group("/api/v1/tasks")
	tasks.Get("/", taskHandler.Index)
	tasks.Post("/", taskHandler.Create)

	tasks.Get("/:id", taskHandler.Show)      // GET /api/v1/tasks/:id
	tasks.Put("/:id", taskHandler.Update)    // PUT /api/v1/tasks/:id
	tasks.Delete("/:id", taskHandler.Delete) // DELETE /api/v1/tasks/:id

	return app
}

func (app *Application) Listen(addr string) error {
	log.Printf("Server listening on %s", addr)
	return app.FiberApp.Listen(addr)
}
