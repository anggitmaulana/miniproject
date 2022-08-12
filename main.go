package main

import (
	"fmt"
	"log"
	"miniproject/controllers"
	"miniproject/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err.Error())
	}

	noteControler := &controllers.NoteControllers{}

	router := httprouter.New()
	router.GET("/", noteControler.Index)
	router.GET("/create", noteControler.Create)
	router.POST("/create", noteControler.Create)
	router.GET("/edit/:id", noteControler.Edit)

	// fmt.Println("aman boss")
	port := ":8080"
	fmt.Println("Aplikasi ini jalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, router))

}
