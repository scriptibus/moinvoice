package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name string
	Mail string
}

type Project struct {
	gorm.Model
	Name               string
	NetHourlyRateCents int
	Customer           Customer
}

type Booking struct {
	gorm.Model
	DurationQuarterHours int
	Description          string
	Project              Project
}
