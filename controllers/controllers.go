package controllers

import (
	"github.com/gin-gonic/gin"

	"genesis-ses4/services"
	"genesis-ses4/utils"
)

func GetRate(c *gin.Context) {
	curRate, err := services.Rate.GetCurrentRate()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// send plain-text string but mark it as JSON to avoid parsing float and
	// formatting (e.g. with custom marshal-json func) as decimal with 2 digit precision
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.String(200, curRate)
}

func SubscribeEmail(c *gin.Context) {
	email := c.Request.FormValue("email")

	// validate email
	if err := utils.ValidateEmail(email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check if already subscribed
	subscribed, err := services.Subscriptions.IsSubscribed(email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if subscribed {
		c.JSON(409, gin.H{"message": "Already subscribed"})
		return
	}

	// add to db
	_, err = services.Subscriptions.AddSubscriber(email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}

func GetSubs(c *gin.Context) {
	emails, err := services.Subscriptions.GetSubscribersEmails()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, emails)
}
