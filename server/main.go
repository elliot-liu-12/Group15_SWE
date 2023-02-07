package main

//imports commented out to avoid generating errors for unused
/*import (
	//"net/http"
	//"errors"
	//"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
	//"gorm.io/gorm/sqlite"
)*/

type movie struct {
	title   string
	runtime float32
}

type user struct {
	//gorm.Model
	username  string
	password  string
	watchlist []movie
	genres    []string
	rating    float32
	providers []string
}

func main() {

}
