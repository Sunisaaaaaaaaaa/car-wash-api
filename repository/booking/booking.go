package booking

import (
	"booking-api/config"
	"booking-api/models"
	invoiceRepo "booking-api/repository/invoice"
	"fmt"
	"time"
)

type BookingRepository interface {
	CreateBookingRepo(req BookingCreateReq) (BookingRes, error)
	GetBookingByIdRepo(id uint) (BookingRes, error)
	GetBookingByCusIdRepo(cusId uint) ([]BookingRes, error)
	GetBookingByEmpIdRepo(emId uint) ([]BookingRes, error)
	GetBookingByDateRepo(req BookingByDateReq) ([]BookingRes, error)
	GetBookingThatNotTakenByDateRepo(req BookingByDateReq) ([]BookingRes, error)
	UpdateToTakeBookingRepo(req UpdateTakenBookingReq) (BookingRes, error)
	UpdateToFinishBookingRepo(req UpdateFinishedBookingReq) (BookingRes, error)
	DeleteBookingRepo(req CancleBookingReq) error
}

type bookingRepository struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{}
}

type BookingCreateReq struct {
	CustomerId  uint      `json:"customerId"`
	BookingDate time.Time `json:"bookingDate"`
	VehicleId   uint      `json:"vehicleId"`
}

type BookingRes struct {
	BookingId   uint      `json:"bookingId"`
	BookingDate time.Time `json:"bookingDate"`
	FinishDate  time.Time `json:"finishDate"`
	IsFinished  bool      `json:"isFinished"`
	IsTaken     bool      `json:"isTaken"`
	TakenAt     time.Time `json:"takenAt"`
	EmployeeId  *uint     `json:"employeeId"`
	CustomerId  uint      `json:"customerId"`
	VehicleId   uint      `json:"vehicleId"`
}

func (r *bookingRepository) CreateBookingRepo(req BookingCreateReq) (BookingRes, error) {
	var userExists models.Customer
	var vehicleExists models.Vehicle

	if err := config.DB.Table("customers").
		Where("id = ?", req.CustomerId).
		Find(&userExists).Error; err != nil {
		return BookingRes{}, err
	}

	if userExists == (models.Customer{}) {
		var err = fmt.Errorf("user does not exist")
		return BookingRes{}, err
	}

	if err := config.DB.Table("vehicles").
		Where("id = ? AND customer_id = ?", req.VehicleId, req.CustomerId).
		Find(&vehicleExists).Error; err != nil {
		return BookingRes{}, err
	}

	if vehicleExists == (models.Vehicle{}) {
		var err = fmt.Errorf("vehicle does not exist")
		return BookingRes{}, err
	}

	newBooking := models.Booking{
		BookingDate: req.BookingDate,
		CustomerId:  req.CustomerId,
		IsTaken:     false,
		IsFinished:  false,
		VehicleId:   req.VehicleId,
		EmployeeId:  nil,
	}
	if err := config.DB.Table("bookings").Create(&newBooking).Error; err != nil {
		return BookingRes{}, err
	}

	newInvoice := models.Invoice{
		BookingId:  newBooking.ID,
		IsCancled:  false,
		IsPaid:     false,
		CustomerId: req.CustomerId,
		VehicleId:  req.VehicleId,
		EmployeeId: nil,
	}

	_, err := invoiceRepo.NewInvoiceRepository().CreateInvoiceRepo(newInvoice)
	if err != nil {
		return BookingRes{}, err
	}

	res := BookingRes{
		BookingId:   newBooking.ID,
		BookingDate: newBooking.BookingDate,
		FinishDate:  newBooking.FinishDate,
		IsFinished:  newBooking.IsFinished,
		IsTaken:     newBooking.IsTaken,
		TakenAt:     newBooking.TakenAt,
		EmployeeId:  newBooking.EmployeeId,
		CustomerId:  newBooking.CustomerId,
		VehicleId:   newBooking.VehicleId,
	}
	return res, nil

}

func (r *bookingRepository) GetBookingByIdRepo(id uint) (BookingRes, error) {
	var res models.Booking

	if err := config.DB.Table("bookings").
		Where("id = ?", id).Find(&res).Error; err != nil {
		return BookingRes{}, err
	}

	if res == (models.Booking{}) {
		var err = fmt.Errorf("booking does not exist")
		return BookingRes{}, err
	}

	result := BookingRes{
		BookingId:   res.ID,
		BookingDate: res.BookingDate,
		FinishDate:  res.FinishDate,
		IsFinished:  res.IsFinished,
		IsTaken:     res.IsTaken,
		TakenAt:     res.TakenAt,
		EmployeeId:  res.EmployeeId,
		CustomerId:  res.CustomerId,
		VehicleId:   res.VehicleId,
	}

	return result, nil
}

func (r *bookingRepository) GetBookingByCusIdRepo(cusId uint) ([]BookingRes, error) {
	var res []models.Booking

	if err := config.DB.Table("bookings").
		Where("customer_id = ?", cusId).Find(&res).Error; err != nil {
		return []BookingRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("booking does not exist")
		return []BookingRes{}, err
	}

	var result []BookingRes
	for _, b := range res {
		tmp := BookingRes{
			BookingId:   b.ID,
			BookingDate: b.BookingDate,
			FinishDate:  b.FinishDate,
			IsFinished:  b.IsFinished,
			IsTaken:     b.IsTaken,
			TakenAt:     b.TakenAt,
			EmployeeId:  b.EmployeeId,
			CustomerId:  b.CustomerId,
			VehicleId:   b.VehicleId,
		}
		result = append(result, tmp)
	}

	return result, nil

}


func (r *bookingRepository) GetBookingByEmpIdRepo(emId uint) ([]BookingRes, error) {
	var res []models.Booking

	if err := config.DB.Table("bookings").
		Where("employee_id = ? AND is_finished = false", emId).Find(&res).Error; err != nil {
		return []BookingRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("booking does not exist")
		return []BookingRes{}, err
	}

	var result []BookingRes
	for _, b := range res {
		if !b.IsTaken {
			b.IsTaken = true
		}
		tmp := BookingRes{
			BookingId:   b.ID,
			BookingDate: b.BookingDate,
			FinishDate:  b.FinishDate,
			IsFinished:  b.IsFinished,
			IsTaken:     b.IsTaken,
			TakenAt:     b.TakenAt,
			EmployeeId:  b.EmployeeId,
			CustomerId:  b.CustomerId,
			VehicleId:   b.VehicleId,
		}
		result = append(result, tmp)
	}

	return result, nil
}

type BookingByDateReq struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}


func (r *bookingRepository) GetBookingByDateRepo(req BookingByDateReq) ([]BookingRes, error) {
	var res []models.Booking
	if err := config.DB.Table("bookings").
		Where("booking_date BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Find(&res).Error; err != nil {
		return []BookingRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("booking does not exist")
		return []BookingRes{}, err
	}

	var result []BookingRes
	for _, b := range res {
		tmp := BookingRes{
			BookingId:   b.ID,
			BookingDate: b.BookingDate,
			FinishDate:  b.FinishDate,
			IsFinished:  b.IsFinished,
			IsTaken:     b.IsTaken,
			TakenAt:     b.TakenAt,
			EmployeeId:  b.EmployeeId,
			CustomerId:  b.CustomerId,
			VehicleId:   b.VehicleId,
		}
		result = append(result, tmp)
	}

	return result, nil

}


func (r *bookingRepository) GetBookingThatNotTakenByDateRepo(req BookingByDateReq) ([]BookingRes, error) {
	var res []models.Booking
	if err := config.DB.Table("bookings").
		Where("is_taken = false AND booking_date BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Find(&res).Error; err != nil {
		return []BookingRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("booking does not exist")
		return []BookingRes{}, err
	}

	var result []BookingRes
	for _, b := range res {
		tmp := BookingRes{
			BookingId:   b.ID,
			BookingDate: b.BookingDate,
			FinishDate:  b.FinishDate,
			IsFinished:  b.IsFinished,
			IsTaken:     b.IsTaken,
			TakenAt:     b.TakenAt,
			EmployeeId:  b.EmployeeId,
			CustomerId:  b.CustomerId,
			VehicleId:   b.VehicleId,
		}
		result = append(result, tmp)
	}

	return result, nil
}

type UpdateTakenBookingReq struct {
	EmployeeId *uint `json:"employeeId"`
	BookingId  uint  `json:"bookingId"`
}


func (r *bookingRepository) UpdateToTakeBookingRepo(req UpdateTakenBookingReq) (BookingRes, error) {
	var before models.Booking

	if err := config.DB.Table("bookings").
		Where("id = ? AND is_taken = false AND is_finished = false", req.BookingId).Find(&before).Error; err != nil {
		return BookingRes{}, err
	}

	if before == (models.Booking{}) {
		var err = fmt.Errorf("booking does not exist")
		return BookingRes{}, err
	}

	before.IsTaken = true
	before.TakenAt = time.Now()
	before.EmployeeId = req.EmployeeId

	if err := config.DB.Table("bookings").Save(&before).Error; err != nil {
		return BookingRes{}, err
	}

	invoice, err := invoiceRepo.NewInvoiceRepository().GetInvoiceByBookingIdRepo(before.ID)
	if err != nil {
		return BookingRes{}, err
	}

	invoice.EmployeeId = req.EmployeeId
	invoice.IsCancled = false
	if err := config.DB.Table("invoices").Save(&invoice).Error; err != nil {
		return BookingRes{}, err
	}

	result := BookingRes{
		BookingId:   before.ID,
		BookingDate: before.BookingDate,
		FinishDate:  before.FinishDate,
		IsFinished:  before.IsFinished,
		IsTaken:     before.IsTaken,
		TakenAt:     before.TakenAt,
		EmployeeId:  before.EmployeeId,
		CustomerId:  before.CustomerId,
		VehicleId:   before.VehicleId,
	}

	return result, nil

}

type UpdateFinishedBookingReq struct {
	BookingId  uint  `json:"bookingId"`
	EmployeeId *uint `json:"employeeId"`
	TotalPrice int   `json:"totalPrice"`
	Discount   int   `json:"discount"`
}

func (r *bookingRepository) UpdateToFinishBookingRepo(req UpdateFinishedBookingReq) (BookingRes, error) {
	var before models.Booking

	if err := config.DB.Table("bookings").
		Where("id = ? AND employee_id = ? AND is_finished = false", req.BookingId, req.EmployeeId).Find(&before).Error; err != nil {
		return BookingRes{}, err
	}

	if before == (models.Booking{}) {
		var err = fmt.Errorf("booking does not exist")
		return BookingRes{}, err
	}

	before.IsFinished = true
	before.FinishDate = time.Now()
	if err := config.DB.Table("bookings").Save(&before).Error; err != nil {
		return BookingRes{}, err
	}

	invoice := invoiceRepo.UpdateInvoiceReq{
		EmployeeId: *req.EmployeeId,
		BookingId:  req.BookingId,
		TotalPrice: req.TotalPrice,
		Discount:   req.Discount,
	}
	_, err := invoiceRepo.NewInvoiceRepository().UpdateInvoiceRepo(invoice)
	if err != nil {
		return BookingRes{}, err
	}

	result := BookingRes{
		BookingId:   before.ID,
		BookingDate: before.BookingDate,
		FinishDate:  before.FinishDate,
		IsFinished:  before.IsFinished,
		IsTaken:     before.IsTaken,
		TakenAt:     before.TakenAt,
		EmployeeId:  before.EmployeeId,
		CustomerId:  before.CustomerId,
		VehicleId:   before.VehicleId,
	}

	return result, nil

}

type CancleBookingReq struct {
	CustomerId uint
	BookingId  uint `form:"bookingId"`
}

func (r *bookingRepository) DeleteBookingRepo(req CancleBookingReq) error {
	var res models.Booking

	if err := config.DB.Table("bookings").
		Where("id = ? AND customer_id = ? AND is_taken = false", req.BookingId, req.CustomerId).Find(&res).Error; err != nil {
		return err
	}

	if res == (models.Booking{}) {
		var err = fmt.Errorf("booking does not exist")
		return err
	}

	_, err := invoiceRepo.NewInvoiceRepository().UpdateCancleInvoiceRepo(req.BookingId)
	if err != nil {
		return err
	}

	if err := config.DB.Table("bookings").
		Where("id = ? AND customer_id = ? AND is_taken = false", req.BookingId, req.CustomerId).
		Delete(&res).Error; err != nil {
		return err
	}

	return nil

}
