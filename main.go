package main

import (
	"fmt"
	"log"
	"miniproject/config"
	"miniproject/controllers"
	"miniproject/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := config.ConnectionDatabase()
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
	router.POST("/create", noteControler.Store)
	router.GET("/edit/:id", noteControler.Edit)
	router.POST("/edit/:id", noteControler.Update)
	router.POST("/done/:id", noteControler.Done)

	// fmt.Println("aman boss")
	port := ":8080"
	fmt.Println("Aplikasi ini jalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, router))

}
