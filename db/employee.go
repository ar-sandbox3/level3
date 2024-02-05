package db

// CreateEmployee creates a new employee.
func (q *Queries) CreateEmployee(name string, departmentID int64) (int64, error) {
	query := q.db.Rebind(`INSERT INTO employee (name, department_id) VALUES (?, ?)`)
	res, err := q.db.Exec(query, name, departmentID)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// GetEmployeeByID returns an employee by their ID.
func (q *Queries) GetEmployeeByID(id int64) (Employee, error) {
	var i Employee
	query := q.db.Rebind(`SELECT name, department_id FROM employee WHERE id = ?`)
	err := q.db.Get(&i, query, id)
	return i, err
}
