package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/dictionary-api/controllers"
	"github.com/stomas418/dictionary-api/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	h := controllers.NewBaseHandler(db)

	router := gin.Default()
	router.GET("/:letter", h.GetWords)
	router.GET("/:letter/:word", h.GetWord)
	router.Run("localhost:8080")
}
