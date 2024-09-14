package bd_logic

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type nameDB struct {
	name string
}

func Logic_postgres(user, password, dbname string) {
	// Logic_postgres establishes a connection to a PostgreSQL database and performs various CRUD operations.
	// Parameters: user, password, dbname - PostgreSQL connection credentials.
	// Return type: None

	// подключение к бд
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// запрос на добавление
	result, err := db.Exec("insert into first_table (name) values ('G0tem')")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк

	// запрос на все строки
	rows, err := db.Query("select * from first_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	namesDB := []nameDB{}

	for rows.Next() {
		p := nameDB{}
		err := rows.Scan(&p.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		namesDB = append(namesDB, p)
	}
	for _, p := range namesDB {
		fmt.Println(p.name)
	}

	// запрос на 1 строку
	row := db.QueryRow("select * from first_table where name = $1", "Юра")
	nam := nameDB{}
	err = row.Scan(&nam.name)
	if err != nil {
		panic(err)
	}
	fmt.Println(nam.name)

	// обновляем строку с name=G0tem
	nameUpdate, err := db.Exec("update first_table set name = $1 where name = $2", "Вася", "G0tem")
	if err != nil {
		panic(err)
	}
	fmt.Println(nameUpdate.RowsAffected()) // количество обновленных строк

	// удаляем строку с name="Вася"
	nameDel, err := db.Exec("delete from first_table where name = $1", "Вася")
	if err != nil {
		panic(err)
	}
	fmt.Println(nameDel.RowsAffected()) // количество удаленных строк
}
