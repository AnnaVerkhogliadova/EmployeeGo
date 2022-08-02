package restAPITest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage *MemoryStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateEmployee(c *fiber.Ctx) error {
	var employee Employee
	if err := c.BodyParser(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.Error(),
		})
		return c.Next()
	}
	createdEmployee := h.storage.Insert(&employee)
	c.Status(fiber.StatusOK).JSON(createdEmployee)
	return nil
}

func (h *Handler) UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee Employee

	if err := c.BodyParser(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.Error(),
		})
		return c.Next()
	}
	updatedEmployee := h.storage.Update(id, &employee)

	return c.Status(fiber.StatusOK).JSON(updatedEmployee)
}

func (h *Handler) GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	employee, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get employee %s\n", err.Error())
		c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.Error(),
		})
		return c.Next()
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func (h *Handler) DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	h.storage.Delete(id)
	c.Status(fiber.StatusOK).JSON("employee deleted")
	return nil
}

func (h *Handler) AllEmployee(c *fiber.Ctx) error {
	employeeArray := h.storage.All()
	c.Status(fiber.StatusOK).JSON(employeeArray)
	return nil
}
