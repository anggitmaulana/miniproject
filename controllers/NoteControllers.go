package controllers

import (
	"log"
	"miniproject/models"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NoteControllers struct {
}

func (controllers *NoteControllers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/index.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var notes []models.Note
	db.Find(&notes)

	datas := map[string]interface{}{
		"Notes": notes,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}

}

func (controllers *NoteControllers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		note := models.Note{
			Assignee: r.FormValue("assignee"),
			Content:  r.FormValue("content"),
			Date:     r.FormValue("date"),
		}
		result := db.Create(&note)
		if result.Error != nil {
			log.Println(result.Error)
		}
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		files := []string{
			"./views/base.html",
			"./views/create.html",
		}
		htmlTemplate, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		err = htmlTemplate.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
		}
	}

}

func (controllers *NoteControllers) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	files := []string{
		"./views/base.html",
		"./views/edit.html",
	}

	templateHtml, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var note models.Note
	db.Where("ID=?", params.ByName("id")).Find(&note)

	data := map[string]interface{}{
		"Note": note,
	}

	err = templateHtml.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}

}
