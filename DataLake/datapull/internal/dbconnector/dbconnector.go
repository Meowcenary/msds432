package dbconnector

import (
  "database/sql"
  // "errors"
  "fmt"
  "log"
  "strings"

  _ "github.com/lib/pq"
)

// Connection details
var (
  Hostname = "localhost"
  Port     = 5431
  Username = "myuser"
  Password = "mypassword"
  Database = "msds432"
)

func openConnection() (*sql.DB, error) {
  // connection string
  conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    Hostname, Port, Username, Password, Database)

  // open database
  db, err := sql.Open("postgres", conn)
  if err != nil {
    return nil, err
  }
  return db, nil
}

func CountData(tableName string) error {
  var count int

  db, err := openConnection()
  if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
  }

  query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
  err = db.QueryRow(query).Scan(&count)

  switch {
    case err != nil:
        log.Fatal(err)
    default:
        fmt.Printf("Number of rows for %s is %d\n", tableName, count)
  }

  return nil
}

// Insert data into specified table.
func InsertData(tableName string, data map[string]interface{}) error {
  db, err := openConnection()
  if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
  }

	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))
	placeholders := make([]string, 0, len(data))

	i := 1
	for column, value := range data {
		columns = append(columns, column)
		values = append(values, value)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		i++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	_, err = db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}
