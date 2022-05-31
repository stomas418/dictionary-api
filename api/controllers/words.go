package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stomas418/dictionary-api/models"
)

func (h *BaseHandler) GetWord(c *gin.Context) {
	wordToSearch := c.Param("word")
	var databaseWord models.DatabaseWord
	query := fmt.Sprintf("SELECT * FROM %s WHERE word = %q", c.Param("letter"), wordToSearch)
	if row, err := h.db.Query(query); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Word not found."})
		return
	} else {
		for row.Next() {
			if err := row.Scan(&databaseWord.Word, &databaseWord.Meanings); err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Database error: %s", err)})
				return
			}
		}
	}

	c.IndentedJSON(http.StatusOK, databaseWord)
}

func getPageNumber(page string) int {
	if page == "" {
		return 1
	} else {
		page_number, err := strconv.Atoi(page)
		if err != nil {
			return 1
		}
		return page_number
	}
}
func (h *BaseHandler) GetWords(c *gin.Context) {
	letter := c.Param("letter")
	page := c.Query("page")
	page_number := getPageNumber(page)

	var databaseWord models.DatabaseWord
	var databaseWords []models.DatabaseWord
	page_word_limit := 100
	lower_limit := page_word_limit * (page_number - 1)
	upper_limit := lower_limit + page_word_limit

	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d,%d", letter, lower_limit, upper_limit)
	if row, err := h.db.Query(query); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("No words starting with letter %s", letter)})
		return
	} else {
		for row.Next() {
			if err := row.Scan(&databaseWord.Word, &databaseWord.Meanings); err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Database error: %s", err)})
				return
			}
			databaseWords = append(databaseWords, databaseWord)
		}
	}

	c.IndentedJSON(http.StatusOK, databaseWords)
}
