package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scriptibus/moinvoice/internal/models"
	"github.com/scriptibus/moinvoice/internal/views"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BookingHandler struct {
	db  *gorm.DB
	log *zap.SugaredLogger
}

func NewBookingHandler(db *gorm.DB, log *zap.SugaredLogger) *BookingHandler {
	return &BookingHandler{db, log}
}

func (b *BookingHandler) BookingList(ctx *gin.Context) {
	var bookings []models.Booking
	b.db.Find(&bookings)

	ctx.HTML(http.StatusOK, "", views.BookingList(bookings))
}

func (b *BookingHandler) BookingFormGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", views.BookingForm())
}

func (b *BookingHandler) BookingFormPost(ctx *gin.Context) {
	// Retrieve form data
	duration, err := strconv.Atoi(ctx.PostForm("duration"))
	if err != nil {
		b.log.Errorw("error converting duration", "error", err)
		ctx.HTML(http.StatusBadRequest, "", views.ErrorPage("Invalid duration format"))
		return
	}

	// Map form data to Booking struct
	booking := models.Booking{
		DurationQuarterHours: duration, // assuming QuarterHour is a custom type that accepts a string
	}

	b.db.Create(&booking)
	ctx.Redirect(http.StatusFound, "/booking/list")
}
