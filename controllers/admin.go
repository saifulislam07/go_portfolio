package controllers

import (
	"fmt"
	"hello_gin/models"
	"net/http"
	"os"
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

	c.Redirect(http.StatusSeeOther, "/interestList")
}

func InterestEdit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
	}

	interest := models.InterestFind(id)

	c.HTML(http.StatusOK,
		"admin/interestedit.html",
		gin.H{
			"title":    "Interest Edit",
			"interest": interest,
		},
	)
}

func InterestUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	existingInterest := models.InterestFind(id)
	if existingInterest == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	updatedInterest := c.PostForm("interest")

	existingInterest.Interest = updatedInterest

	if err := models.DB.Save(existingInterest).Error; err != nil {
		fmt.Printf("Error saving interest: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/interestList")
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

	c.Redirect(http.StatusSeeOther, "/interestList")
}

func PortfolioList(c *gin.Context) {
	var portfolios []models.Portfolios

	// Fetch all portfolios (newest first)
	if err := models.DB.Order("id desc").Find(&portfolios).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error fetching portfolios: %v", err))
		return
	}

	// Render template with data
	c.HTML(http.StatusOK, "admin/portfolioList.html", gin.H{
		"title":      "Portfolios",
		"portfolios": portfolios,
	})
}

func Portfolios(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/portfolios.html", gin.H{
		"title": "Portfolios",
	})
}

func PortfoliosCreate(c *gin.Context) {
	var portfolio models.Portfolios

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "No image uploaded")
		return
	}

	// Save uploaded file
	filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving file: %v", err))
		return
	}

	// Always create a new portfolio record
	portfolio.Image = filename
	if err := models.DB.Create(&portfolio).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating portfolio: %v", err))
		return
	}

	c.Redirect(http.StatusSeeOther, "/portfolioList")
}

func PortfoliosEdit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
	}

	portfolio := models.PortfolioFind(id)

	c.HTML(http.StatusOK,
		"admin/portfolioedit.html",
		gin.H{
			"title":     "Portfolio Edit",
			"portfolio": portfolio,
		},
	)
}

func PortfoliosUpdate(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolios

	// Find the existing portfolio
	if err := models.DB.First(&portfolio, id).Error; err != nil {
		c.String(http.StatusNotFound, "Portfolio not found")
		return
	}

	// Handle new image upload (optional)
	file, err := c.FormFile("image")
	if err == nil {
		// Save new image file
		filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving file: %v", err))
			return
		}

		// Optional: delete old image file if it exists
		if portfolio.Image != "" {
			_ = os.Remove(portfolio.Image)
		}

		portfolio.Image = filename
	}

	// Save updated portfolio
	if err := models.DB.Save(&portfolio).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error updating portfolio: %v", err))
		return
	}

	c.Redirect(http.StatusSeeOther, "/portfolioList")
}

func PortfoliosDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error converting id: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	portfolio := models.PortfolioFind(id)
	if portfolio == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := models.DB.Delete(portfolio).Error; err != nil {
		fmt.Printf("Error deleting portfolio: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/portfolioList")
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
