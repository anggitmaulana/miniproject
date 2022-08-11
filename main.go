package main

import (
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

	// fmt.Println("aman boss")
	log.Fatal(http.ListenAndServe(":8080", router))

}
