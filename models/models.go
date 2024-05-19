package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Role      int
}

type Invoice struct {
	gorm.Model
	TotalPrice int
	Discount   int
	Amount     int
	IsCancled  bool
	IsPaid     bool
	EndDate    time.Time
	CustomerId uint
	Customer   Customer `gorm:"references:ID"`
	BookingId  uint
	Booking    Booking `gorm:"references:ID"`
	VehicleId  uint
	Vehicle    Vehicle `gorm:"references:ID"`
	EmployeeId *uint
	Employee   Employee `gorm:"references:ID"`
}

type Vehicle struct {
	gorm.Model
	Size       string
	Brand      string
	Color      string
	RegisNo    string
	CustomerId uint
	Customer   Customer `gorm:"references:ID"`
}

type Booking struct {
	gorm.Model
	BookingDate time.Time
	FinishDate  time.Time
	IsFinished  bool
	IsTaken     bool
	TakenAt     time.Time
	EmployeeId  *uint
	Employee    Employee `gorm:"references:ID"`
	CustomerId  uint
	Customer    Customer `gorm:"references:ID"`
	VehicleId   uint
	Vehicle     Vehicle `gorm:"references:ID"`
}

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Role      int
	IsActive  bool
	Gender    string
	Age       string
	Image     string
}
