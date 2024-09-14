package bd_logic_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	user     string
	password string
	dbname   string
}

type nameDB struct {
	name string
}

func getDBConfig() dbConfig {
	// получить текущий путь
	currentDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	// получить путь к родительской директории
	parentDir := filepath.Dir(filepath.Dir(currentDir))

	// загрузить переменные из .env файла
	err = godotenv.Load(parentDir + "/.env")

	// err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return dbConfig{
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB_TEST"),
	}
}

func TestLogic_postgres(t *testing.T) {
	dbConfig := getDBConfig()

	// Test successful connection to PostgreSQL database
	t.Run("test successful connection", func(t *testing.T) {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
	})

	// Test successful insertion of data into `first_table`
	t.Run("test successful insertion", func(t *testing.T) {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
		result, err := db.Exec("insert into first_table (name) values ('test_name')")
		if err != nil {
			t.Fatal(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}
		if rowsAffected != 1 {
			t.Errorf("expected 1 row affected, got %d", rowsAffected)
		}
	})

	// Test successful retrieval of all rows from `first_table`
	t.Run("test successful retrieval of all rows", func(t *testing.T) {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query("select * from first_table")
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()
		var namesDB []nameDB
		for rows.Next() {
			p := nameDB{}
			err := rows.Scan(&p.name)
			if err != nil {
				t.Fatal(err)
			}
			namesDB = append(namesDB, p)
		}
		if len(namesDB) != 1 {
			t.Errorf("expected 1 row, got %d", len(namesDB))
		}
	})

	// Test successful retrieval of a single row from `first_table`
	t.Run("test successful retrieval of a single row", func(t *testing.T) {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
		row := db.QueryRow("select * from first_table where name = $1", "test_name")
		nam := nameDB{}
		err = row.Scan(&nam.name)
		if err != nil {
			t.Fatal(err)
		}
		if nam.name != "test_name" {
			t.Errorf("expected name 'test_name', got '%s'", nam.name)
		}
	})

	// Test successful update of a row in `first_table`
	t.Run("test successful update", func(t *testing.T) {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
		nameUpdate, err := db.Exec("update first_table set name = $1 where name = $2", "new_name", "test_name")
		if err != nil {
			t.Fatal(err)
		}
		rowsAffected, err := nameUpdate.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}
		if rowsAffected != 1 {
			t.Errorf("expected 1 row affected, got %d", rowsAffected)
		}
	})

	// Test successful deletion of a row from `first_table`
	t.Run("test successful deletion", func(t *testing.T) {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
		nameDel, err := db.Exec("delete from first_table where name = $1", "new_name")
		if err != nil {
			t.Fatal(err)
		}
		rowsAffected, err := nameDel.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}
		if rowsAffected != 1 {
			t.Errorf("expected 1 row affected, got %d", rowsAffected)
		}
	})

	// Test error handling for invalid SQL queries
	t.Run("test error handling for invalid SQL queries", func(t *testing.T) {
		connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
		_, err = db.Exec("invalid SQL query")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	// // Test error handling for invalid database credentials
	// t.Run("test error handling for invalid database credentials", func(t *testing.T) {
	// 	user := "invalid_user"
	// 	password := "invalid_password"
	// 	dbname := "invalid_db"
	// 	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	// 	_, err := sql.Open("postgres", connStr)
	// 	if err == nil {
	// 		t.Errorf("expected error, got nil")
	// 	}
	// })
}
