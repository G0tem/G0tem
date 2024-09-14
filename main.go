package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	bot "github.com/G0tem/G0tem/bot"
	house "github.com/G0tem/G0tem/src/house_logic"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type nameDB struct {
	name string
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Функция логики работы с бд
	logic_DB()

	// логика бота
	go bot.RunBot()
	fmt.Println(house.MyHouse())
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func logic_DB() {
	// загрузка переменных
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке переменных окружения из файла .env")
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

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

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
