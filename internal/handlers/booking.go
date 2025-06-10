package handlers

import (
	"net/http"

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
