package main

import (
	"fmt"
	"my-fiber-app/handles"
	"my-fiber-app/repositories"

	// "ariga.io/atlas-provider-gorm/gormschema"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// run this code after with lib schema loader 
	// stmts, err := gormschema.LoadGormSchema()
    // if err != nil {
    //     fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
    //     os.Exit(1)
    // }

    // gormschema.PrintSchema(stmts)

	// the first one to run make migrations folder
	// stmts, err := gormschema.New("mysql").Load(&entities.Task{})
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
	// 	os.Exit(1)
	// }
	// io.WriteString(os.Stdout, stmts)

	dsn := "root:pass@tcp(localhost:3306)/crud?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if db != nil {
		fmt.Println("connection success")
	}

	if err != nil {
		panic("failed to connect to database")
	}

	taskRepository := repositories.NewTaskRepository(db)

	taskHandler := handles.NewTaskHandle(taskRepository)

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Worlds!")
	})

	app.Post("/task", taskHandler.CreateTask)
	app.Get("/task", taskHandler.GetTasks)
	app.Get("/task/:id", taskHandler.GetTasksById)
	app.Put("/task/:id", taskHandler.UpdateTask)
	app.Delete("/task/:id", taskHandler.DeleteTasksById)

	app.Listen(":3000")
}
