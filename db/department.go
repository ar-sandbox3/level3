package db

// CreateDepartment creates a new department.
func (q *Queries) CreateDepartment(name string) (int64, error) {
	query := q.db.Rebind(`INSERT INTO department (name) VALUES (?)`)
	res, err := q.db.Exec(query, name)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// GetDepartmentByID returns a department by its ID.
func (q *Queries) GetDepartmentByID(id int64) (Department, error) {
	var i Department
	query := q.db.Rebind(`SELECT * FROM department WHERE id = ?`)
	err := q.db.Get(&i, query, id)
	return i, err
}
