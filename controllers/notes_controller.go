package controllers

import (
	"fmt"
	"hello_gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()

	c.HTML(http.StatusOK, "notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(http.StatusOK, "notes/new.html",
		gin.H{},
	)
}

func NotesCreate(c *gin.Context) {
	name := c.PostForm("name")
	content := c.PostForm("content")

	models.NotesCreate(name, content)

	c.Redirect(http.StatusSeeOther, "/notes")
}

func NotesShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	note := models.NotesFind(id)

	c.HTML(http.StatusOK,
		"notes/show.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesEdit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
	}

	note := models.NotesFind(id)

	c.HTML(http.StatusOK,
		"notes/edit.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	note := models.NotesFind(id)
	if note == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	name := c.PostForm("name")
	content := c.PostForm("content")

	note.Name = name
	note.Content = content

	// *** Save the changes to the database ***
	if err := models.DB.Save(note).Error; err != nil {
		fmt.Printf("Error saving note: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/notes/"+idStr)
}

func NotesDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	note := models.NotesFind(id)
	if note == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := models.DB.Delete(&note).Error; err != nil {
		fmt.Printf("Error deleting note: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Redirect to notes list after deleting
	c.Redirect(http.StatusSeeOther, "/notes")
}
