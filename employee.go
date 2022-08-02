package restAPITest

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name" validate:"required,min=3,max=15"`
	Sex    string `json:"sex"  validate:"required"`
	Age    int    `json:"age" validate:"required,numeric,min=10"`
	Salary int    `json:"salary" validate:"required,numeric,min=0"`
}
