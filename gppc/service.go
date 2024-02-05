package gppc

import (
	"context"
	"fmt"
	"http/employee"
	//"employee/employee.proto"
)

type server struct{}

func (s *server) Create(ctx context.Context, req *employee.CreateRequest) (*employee.CreateResponse, error) {
	// Ваша логика создания сотрудника
	fmt.Printf("Received CreateRequest: %+v\n", req)

	// Возвращаем заглушенный CreateResponse
	return &employee.CreateResponse{
		Id:     req.Id,
		Name:   req.Name,
		Sex:    req.Sex,
		Age:    req.Age,
		Salary: req.Salary,
	}, nil
}
