package main

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
	user     = "postgres"
	password = 1234
	dbname   = "usersGo"
)

var Validator = validator.New()

type IError struct {
	Field string
	Tag   string
	Value string
}

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

func main() {
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

	//router := gin.Default()

	//router.POST("/employee", handler.CreateEmployee)
	//router.GET("/employee/:id", handler.GetEmployee)
	//router.PUT("/employee/:id", handler.UpdateEmployee)
	//router.DELETE("/employee/:id", handler.DeleteEmployee)
	//router.GET("/employee", handler.AllEmployee)

	//router.Run()
}
