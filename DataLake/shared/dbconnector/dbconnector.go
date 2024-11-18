package dbconnector

import (
  "database/sql"
  "fmt"
  // "log"
  "reflect"
  "strings"

  _ "github.com/lib/pq"
)

// Global db instance for the connection pool
var db *sql.DB

// Connection details
var (
  Hostname = "localhost"
  Port     = 5431
  Username = "myuser"
  Password = "mypassword"
  Database = "msds432"
)

// Initialize and open the database connection pool once
func InitDB() error {
  conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    Hostname, Port, Username, Password, Database)

  var err error
  db, err = sql.Open("postgres", conn)
  if err != nil {
    return fmt.Errorf("failed to open database connection: %w", err)
  }

  // Set connection pool limits
  db.SetMaxOpenConns(10)  // adjust as needed
  db.SetMaxIdleConns(5)   // adjust as needed
  return nil
}

// CloseDB closes the database connection pool
func CloseDB() error {
  return db.Close()
}

func CountData(tableName string) error {
  var count int
  query := fmt.Sprintf("SELECT COUNT(*) FROM \"%s\"", tableName)
  err := db.QueryRow(query).Scan(&count)
  if err != nil {
    return fmt.Errorf("failed to count data: %w", err)
  }
  fmt.Printf("Number of rows for %s is %d\n", tableName, count)
  return nil
}

func InsertData[T any](tableName string, data T) error {
  // Using reflection to get the struct fields and values
  val := reflect.ValueOf(data)
  typ := reflect.TypeOf(data)

  // Ensure it's a struct
  if val.Kind() != reflect.Struct {
    return fmt.Errorf("expected struct, got %s", val.Kind())
  }

  // Prepare slices for columns, values, and placeholders
  columns := make([]string, 0)
  values := make([]interface{}, 0)
  placeholders := make([]string, 0)

  // Loop through the struct fields
  for i := 0; i < val.NumField(); i++ {
    field := val.Field(i)
    fieldType := typ.Field(i)
    columnName := fieldType.Name // assuming DB column name matches field name

    // Add field name and value to the lists
    columns = append(columns, columnName)
    values = append(values, field.Interface())
    placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
  }

  // Construct the SQL insert query
  query := fmt.Sprintf("INSERT INTO \"%s\" (%s) VALUES (%s)",
    tableName,
    strings.Join(columns, ", "),
    strings.Join(placeholders, ", "))

  // Execute the query
  _, err := db.Exec(query, values...)
  if err != nil {
    return fmt.Errorf("failed to insert data: %w", err)
  }

  return nil
}