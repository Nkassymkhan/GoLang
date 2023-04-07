package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Cost        float64
}

var db *gorm.DB
var err error

func initDB() {
	db, err = gorm.Open(sqlite.Open("C:\\Users\\77079\\Desktop\\Programming\\Go\assignment3\book_store.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

}

func main() {
	initDB()

}
