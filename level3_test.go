package level3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ar-sandbox3/level3"
	"github.com/ar-sandbox3/level3/db"
)

func TestCreate(t *testing.T) {
	// Initialize the database.
	conn, err := level3.Initialize()
	require.NoError(t, err)

	// Create a new department.
	queries := db.New(conn)
	id, err := queries.CreateDepartment("Engineering")
	require.NoError(t, err)

	// Get the department by ID.
	d, err := queries.GetDepartmentByID(id)
	require.NoError(t, err)

	require.Equal(t, "Engineering", d.Name)

	// Create a new employee.
	id, err = queries.CreateEmployee("Kurnia Abu Hasna", id)
	require.NoError(t, err)

	// Get the employee by ID.
	e, err := queries.GetEmployeeByID(id)
	require.NoError(t, err)
	require.Equal(t, "Kurnia Abu Hasna", e.Name)
	require.Equal(t, d.ID, e.DepartmentID)
}
