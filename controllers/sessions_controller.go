package controllers

import (
	"hello_gin/helpers"
	"hello_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}

func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	available := models.UserAvailability(email)

	if !available {

		c.HTML(
			http.StatusIMUsed,
			"home/signup.html",
			gin.H{
				"alert": "Email already exists",
			},
		)
		return
	}
	if password != confirmPassword {
		c.HTML(
			http.StatusIMUsed,
			"home/signup.html",
			gin.H{
				"alert": "Passwords do not match",
			},
		)
		return
	}

	user := models.UserCreate(email, password)
	if user.ID == 0 {
		c.HTML(
			http.StatusInternalServerError,
			"home/signup.html",
			gin.H{
				"alert": "Unable to create user",
			},
		)
		return
	} else {
		//
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := models.UserCheck(email, password)

	if user.ID == 0 {
		c.HTML(
			http.StatusOK,
			"home/login.html",
			gin.H{
				"alert": "Invalid credentials",
			},
		)
		return
	} else {
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Logout(c *gin.Context) {
	helpers.SessionClear(c)

	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{
			"alert": "Logged out",
		},
	)

	c.Redirect(http.StatusMovedPermanently, "/")
}
