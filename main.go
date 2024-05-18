package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"

	"genesis-ses4/controllers"
	"genesis-ses4/models"
	"genesis-ses4/services"
)

func main() {
	godotenv.Load(".env")
	devMode := os.Getenv("MODE") == "dev"

	var (
		crons              = cron.New()
		updateRateSchedule = "@every 6h"
		sendEmailsSchedule = "0 9 * * *"
	)

	if devMode {
		updateRateSchedule = "@every 7s"
		sendEmailsSchedule = "@every 13s"
	}

	//
	// DB
	//
	log.Println("connecting to db")
	if err := models.InitDB(&models.DBConfig{
		Host:          os.Getenv("DB_HOST"),
		Port:          os.Getenv("DB_PORT"),
		User:          os.Getenv("DB_USER"),
		Pass:          os.Getenv("DB_PASS"),
		Name:          os.Getenv("DB_NAME"),
		Retries:       10,
		RetryInterval: time.Second,
	}); err != nil {
		log.Fatal(err)
	}

	//
	// Rate
	//
	log.Println("setting up rate service")
	if err := services.Rate.Init(os.Getenv("CURRENCY_API_URL")); err != nil {
		log.Fatal(err)
	}
	crons.AddFunc(updateRateSchedule, func() { services.Rate.GetUpdatedRate() })

	//
	// Emails
	//
	log.Println("setting up emails service")
	services.Emails.Init(&services.EmailsConfig{
		Host:    os.Getenv("SMTP_HOST"),
		Port:    os.Getenv("SMTP_PORT"),
		User:    os.Getenv("SMTP_USER"),
		Pass:    os.Getenv("SMTP_PASS"),
		Devmode: devMode,
	})
	sendEmails := services.Emails.NewList(&services.MailingList{
		Subject:       "Today's USD-UAH exchange rate",
		GetRecipients: services.Subscriptions.GetSubscribersEmails,
		GetBody: func() (string, error) {
			rate, err := services.Rate.GetUpdatedRate()
			return fmt.Sprintf("Today's USD to UAH exchange rate: %s", rate), err
		},
	})
	crons.AddFunc(sendEmailsSchedule, sendEmails)

	// Schedule tasks
	crons.Start()

	//
	// API
	//
	log.Println("starting API server")
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		// prevent CORS errors
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})

	r.GET("/rate", controllers.GetRate)
	r.GET("/subs", controllers.GetSubs)
	r.POST("/subscribe", controllers.SubscribeEmail)

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
