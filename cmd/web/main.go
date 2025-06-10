package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scriptibus/moinvoice/internal/handlers"
	"github.com/scriptibus/moinvoice/internal/setup"
)

func main() {
	log, err := setup.InitLogger()
	if err != nil {
		panic("Failed to init logger: " + err.Error())
	}
	log.Info("Initialized logger")

	log.Info("Initializing ORM")
	db, err := setup.InitORM()
	if err != nil {
		panic("Failed to init ORM: " + err.Error())
	}
	log.Info("Initialized ORM")

	log.Info("Initializing Webserver")
	router := gin.New()
	router.Use(setup.GinLoggerHandler(log))
	router.Use(gin.Recovery())
	router.HTMLRender = &setup.TemplRender{}

	bookingHandler := handlers.NewBookingHandler(db, log)
	router.GET("/booking/list", bookingHandler.BookingList)
	router.GET("/booking/create", bookingHandler.BookingFormGet)
	router.POST("/booking/create", bookingHandler.BookingFormPost)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to run webserver: " + err.Error())
	}
}
