package main

import (
	"Quiz-3/controllers"
	"Quiz-3/database"
	"Quiz-3/middleware"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/lib/pq"
)

const ()

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGNAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// Router GIN
	router := gin.Default()
	//endpoint bangun-datar
	bangunDatar := router.Group("/bangun-datar")
	bangunDatar.GET("/segitiga-sama-sisi", controllers.SegitigaSamasisi)
	bangunDatar.GET("/persegi", controllers.Persegi)
	bangunDatar.GET("/persegi-panjang", controllers.PersegiPanjang)
	bangunDatar.GET("/lingkaran", controllers.Lingkaran)

	//endpoint kategori
	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", gin.BasicAuth(middleware.Auth()), controllers.InsertCategory)
	router.PUT("/categories/:id", gin.BasicAuth(middleware.Auth()), controllers.UpdateCategory)
	router.DELETE("/categories/:id", gin.BasicAuth(middleware.Auth()), controllers.DeleteCategory)
	router.GET("/categories/:id/books", controllers.GetBookbyIdCategory)

	//endpoint buku
	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", gin.BasicAuth(middleware.Auth()), controllers.InsertBook)
	router.PUT("/books/:id", gin.BasicAuth(middleware.Auth()), controllers.UpdateBook)
	router.DELETE("/books/:id", gin.BasicAuth(middleware.Auth()), controllers.DeleteBook)
	router.Run(":" + os.Getenv("PORT"))
}
