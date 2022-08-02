package restAPITest

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/lithammer/shortuuid"
)

var (
	id     string
	name   string
	sex    string
	age    int
	salary int
)

type Storage interface {
	Insert(e *Employee) Employee
	Get(id string) (Employee, error)
	Update(id string, e *Employee) Employee
	Delete(id string)
	All() []Employee
}

type MemoryStorage struct {
	db *sql.DB
}

func NewMemoryStorage(Db *sql.DB) *MemoryStorage {
	return &MemoryStorage{
		db: Db,
	}
}

func (s *MemoryStorage) Insert(e *Employee) Employee {
	sqlStatement := "INSERT INTO users (id, name, sex, age, salary) VALUES ($1, $2, $3, $4, $5)"
	var id = shortuuid.New()
	_, err := s.db.Query(sqlStatement, id, e.Name, e.Sex, e.Age, e.Salary)
	if err != nil {
		panic(err)
	}

	return Employee{id, e.Name, e.Sex, e.Age, e.Salary}
}

func (s *MemoryStorage) Get(id string) (Employee, error) {
	result, err := s.db.Query("SELECT * FROM users WHERE id = $1 ", id)

	if err != nil {
		panic(err.Error())
	}
	var emp Employee
	for result.Next() {
		err := result.Scan(&emp.ID, &emp.Name, &emp.Sex, &emp.Age, &emp.Salary)
		if err != nil {
			panic(err.Error())
		}
	}
	return emp, nil
}

func (s *MemoryStorage) Update(id string, e *Employee) Employee {
	sqlStatement := "UPDATE users SET name = $2, sex = $3, age = $4, salary = $5 " +
		"WHERE id = $1 RETURNING id, name, sex, age, salary"
	result, err := s.db.Query(sqlStatement, id, e.Name, e.Sex, e.Age, e.Salary)

	if err != nil {
		panic(err.Error())
	}

	var emp Employee
	for result.Next() {
		err := result.Scan(&emp.ID, &emp.Name, &emp.Sex, &emp.Age, &emp.Salary)
		if err != nil {
			panic(err.Error())
		}
	}
	return emp
}

func (s *MemoryStorage) Delete(id string) {
	st, err := s.db.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		panic(err)
	}
	st.Exec(id)
}

func (s *MemoryStorage) All() []Employee {
	var emp []Employee

	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &sex, &age, &salary)
		if err != nil {
			panic(err)
		}
		emp = append(emp, Employee{id, name, sex, age, salary})
	}
	return emp
}
