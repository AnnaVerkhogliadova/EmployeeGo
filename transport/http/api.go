package http

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	restAPITest "http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = 1234
	dbname   = "employeeGo"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func validate(c *fiber.Ctx) error {
	var errors []*IError
	body := new(restAPITest.Employee)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}

func Api() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	memoryStorage := restAPITest.NewMemoryStorage(Db)
	handler := restAPITest.NewHandler(memoryStorage)

	app := fiber.New()

	app.Post("/employee", validate, handler.CreateEmployee)
	app.Get("/employee/:id", handler.GetEmployee)
	app.Put("/employee/:id", validate, handler.UpdateEmployee)
	app.Get("/employee", handler.AllEmployee)
	app.Delete("/employee/:id", handler.DeleteEmployee)
	app.Listen(":8080")
}
