package main

import (
	"go-restful-api/book"
	"go-restful-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:teuingatuh@tcp(127.0.0.1:8889)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// GO layer
	// main
	// handler
	// service
	// repository
	// db
	// mysql

	// CRUD

	/*CREATE*/
	// book := &book.Book{}
	// book.Title = "Power Ranger"
	// book.Price = 90000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "lorem Ipsum, Lorem Ipsum"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("================================================")
	// 	fmt.Println("Error creating book record: ", err)
	// 	fmt.Println("================================================")
	// }

	/*GET ALL*/
	// var books []book.Book

	// err = db.Debug().Where("title = ?", "Power Ranger").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("================================================")
	// 	fmt.Println("Error finding book record: ", err)
	// 	fmt.Println("================================================")
	// }

	// for _, book := range books {
	// 	fmt.Println("Title: ", book.Title)
	// }

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("================================================")
	// 	fmt.Println("Error finding book record: ", err)
	// 	fmt.Println("================================================")
	// }

	/*DELETE*/
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("================================================")
	// 	fmt.Println("Error deleted book record: ", err)
	// 	fmt.Println("================================================")
	// }

	/*UPDATE*/
	// book.Title = "Mantap"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("================================================")
	// 	fmt.Println("Error updated book record: ", err)
	// 	fmt.Println("================================================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetAllBook)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.Run(":8080");
}
