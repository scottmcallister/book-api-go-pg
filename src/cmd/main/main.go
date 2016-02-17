package main

import (
	"os"
	"database/sql"
	"gopkg.in/gorp.v1"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Book struct {
	Id        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Title string `db:"title" json:"title"`
	Author  string `db:"author" json:"author"`
	Publisher string `db:"publisher" json:"publisher"`
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres", os.Getenv("HEROKU_POSTGRESQL_MAUVE_URL"))
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Book{}, "Book").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.GET("/books", GetBooks)
		v1.GET("/books/:id", GetBook)
		v1.POST("/books", PostBook)
		v1.PUT("/books/:id", UpdateBook)
		v1.DELETE("/books/:id", DeleteBook)
		v1.OPTIONS("/books", OptionsBook)
		v1.OPTIONS("/books/:id", OptionsBook)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

func GetBooks(c *gin.Context) {
	var books []Book
	_, err := dbmap.Select(&books, "SELECT * FROM book")

	if err == nil {
		c.JSON(200, books)
	} else {
		c.JSON(404, gin.H{"error": "no book(s) in the table"})
	}
}

func GetBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Book
	err := dbmap.SelectOne(&book, "SELECT * FROM book WHERE id=$1 LIMIT 1", id)

	if err == nil {
		book_id, _ := strconv.ParseInt(id, 0, 64)

		content := &Book{
			Id:        	book_id,
			Title: 		book.Title,
			Author:  	book.Author,
			Publisher: 	book.Publisher,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "book not found"})
	}
}

func PostBook(c *gin.Context) {
	var book Book
	c.Bind(&book)

	log.Println(book)

	if book.Title != "" && book.Author != "" && book.Publisher != ""{

		if insert, _ := dbmap.Exec(`INSERT INTO book (title, author, publisher) VALUES ($1, $1, $1)`, book.Title, book.Author, book.Publisher); insert != nil {
			book_id, err := insert.LastInsertId()
			if err == nil {
				content := &Book{
					Id:        	book_id,
					Title: 		book.Title,
					Author:  	book.Author,
					Publisher: 	book.Publisher,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
				c.JSON(400, gin.H{"error": "Insert failed"})
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}

func UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Book
	err := dbmap.SelectOne(&book, "SELECT * FROM book WHERE id=$1", id)

	if err == nil {
		var json Book
		c.Bind(&json)

		book_id, _ := strconv.ParseInt(id, 0, 64)

		book := Book{
			Id:        	book_id,
			Title: 		book.Title,
			Author:  	book.Author,
			Publisher: 	book.Publisher,
		}

		if book.Title != "" && book.Author != "" && book.Publisher != "" {
			_, err = dbmap.Update(&book)

			if err == nil {
				c.JSON(200, book)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "book not found"})
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")

	var book Book
	err := dbmap.SelectOne(&book, "SELECT * FROM book WHERE id=$1", id)

	if err == nil {
		_, err = dbmap.Delete(&book)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "book not found"})
	}
}

func OptionsBook(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}