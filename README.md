# level3

This is a way to solve the level3 challenge, at several few commits, this provides a starter template for dealing with basic operations.

To use this code, you need to change the DB connection string used in [level3.go](./level3.go).

In the `Initialize` function:

```go
// Initialize creates a new connection to the database and runs the schema.
func Initialize() (*sqlx.DB, error) {
	// Connect to the database. Please change this to your own database.
	db, err := sqlx.Connect("mysql", "root:password@tcp(localhost:3306)/level3?multiStatements=true")
```

> [!TIP]
> You can run a MySQL server locally using docker.
>
> For example, binding to the host network:
>
> ```console
> docker run --rm -it --net=host -v `pwd`/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=level3 mysql:8
> ```
