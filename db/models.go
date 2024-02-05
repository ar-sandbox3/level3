package db

// Department represents a department.
type Department struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// Employee represents an employee.
type Employee struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	DepartmentID int64  `db:"department_id"`
}
