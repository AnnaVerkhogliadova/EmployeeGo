package main

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	restAPITest "http"
	"http/employee"
	"log"
	"net"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = 1234
	dbname   = "employeeGo"
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

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	employee.RegisterEmployeeServiceServer(s, &server{})

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
