package controllers

import (
	"fmt"
	"hello_gin/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/dashboard.html", gin.H{
		"title": "Dashboard",
	})
}

func About(c *gin.Context) {
	about, err := models.AboutFindLatest()
	if err != nil {
		c.String(http.StatusNotFound, "No about data found")
		return
	}

	c.HTML(http.StatusOK, "admin/about.html", gin.H{
		"title":      "About",
		"AboutMe":    about.AboutMe,
		"AboutImage": about.AboutImage,
	})
}

func AboutCreate(c *gin.Context) {
	aboutMe := c.PostForm("about_me")

	var about models.About
	result := models.DB.Order("id desc").First(&about)

	file, err := c.FormFile("about_image")
	if err == nil {
		filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving file: %v", err))
			return
		}
		about.AboutImage = filename
	}

	about.AboutMe = aboutMe

	if result.RowsAffected == 0 {
		models.DB.Create(&about)
	} else {
		models.DB.Save(&about)
	}

	c.Redirect(http.StatusSeeOther, "/about/new")
}

func Skils(c *gin.Context) {
	skil, _ := models.SkilsFindLatest()

	c.HTML(http.StatusOK, "admin/skils.html", gin.H{
		"title": "Skills",
		"skils": skil.Skils,
	})
}

func SkilCreate(c *gin.Context) {
	skilsText := c.PostForm("skils")

	var skil models.Skils
	result := models.DB.Order("id desc").First(&skil)

	skil.Skils = skilsText

	if result.RowsAffected == 0 {
		models.DB.Create(&skil)
	} else {
		models.DB.Save(&skil)
	}

	c.Redirect(http.StatusSeeOther, "/skils/new")
}

func Interests(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/interests.html", gin.H{
		"title": "Interests",
	})
}

func InterestList(c *gin.Context) {

	interests, err := models.InterestList()

	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching interests: %v", err)
		return
	}

	c.HTML(http.StatusOK, "admin/interestList.html", gin.H{
		"title":     "Interest List",
		"interests": interests,
	})
}

func InterestCreate(c *gin.Context) {
	interestText := c.PostForm("interest")

	interest := models.Interests{
		Interest: interestText,
	}
	result := models.DB.Create(&interest)
	if result.Error != nil {
		c.String(http.StatusInternalServerError, "Error creating interest: %v", result.Error)
		return
	}

	c.Redirect(http.StatusSeeOther, "/interests/new")
}

func InterestEdit(c *gin.Context) {
}
func InterestUpdate(c *gin.Context) {
}

func InterestDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	interest := models.InterestFind(id)
	if interest == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := models.DB.Delete(interest).Error; err != nil {
		fmt.Printf("Error deleting interest: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/InterestList")
}

func Portfolios(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/portfolios.html", gin.H{
		"title": "Portfolios",
	})
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/contact.html", gin.H{
		"title": "Contact",
	})
}

func SocialMedias(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/social-medias.html", gin.H{
		"title": "Social Medias",
	})
}

func Resume(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/resume.html", gin.H{
		"title": "Resume",
	})
}

func Clients(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/clients.html", gin.H{
		"title": "Clients",
	})
}
