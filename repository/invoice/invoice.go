package invoice

import (
	"booking-api/config"
	"booking-api/models"
	"fmt"
	"time"
)

type InvoiceRepository interface {
	CreateInvoiceRepo(req models.Invoice) (InvoiceRes, error)
	GetInvoiceByIdRepo(id uint) (InvoiceRes, error)
	GetInvoiceByBookingIdRepo(id uint) (models.Invoice, error)
	GetInvoiceByDateRepo(req InvoiceByDateReq) ([]InvoiceRes, error)
	GetPaidInvoiceByDateRepo(req InvoiceByDateReq) ([]InvoiceRes, error)
	GetCancledInvoiceByDateRepo(req InvoiceByDateReq) ([]InvoiceRes, error)
	GetInvoiceByEmpIdAndDateRepo(id uint, req InvoiceByDateReq) ([]InvoiceRes, error)
	GetInvoiceByCusIdRepo(id uint) ([]InvoiceRes, error)
	UpdateInvoiceRepo(req UpdateInvoiceReq) (models.Invoice, error)
	UpdatePaidInvoiceRepo(id uint) (InvoiceRes, error)
	UpdateCancleInvoiceRepo(bookingId uint) (models.Invoice, error)
}

type invoiceRepository struct{}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{}
}

type InvoiceRes struct {
	InvoiceId  uint      `json:"invoiceId"`
	TotalPrice int       `json:"totalPrice"`
	Discount   int       `json:"discount"`
	Amount     int       `json:"amount"`
	IsCancled  bool      `json:"isCancled"`
	IsPaid     bool      `json:"isPaid"`
	EndDate    time.Time `json:"endDate"`
	CustomerId uint      `json:"customerId"`
	EmployeeId uint      `json:"empolyeeId"`
	VehicleId  uint      `json:"vehicleId"`
	BookingId  uint      `json:"bookingId"`
}

// test already
func (r *invoiceRepository) CreateInvoiceRepo(req models.Invoice) (InvoiceRes, error) {
	var bookExists models.Booking

	if err := config.DB.Table("bookings").
		Where("id = ?", req.BookingId).
		Find(&bookExists).Error; err != nil {
		return InvoiceRes{}, err
	}

	if (bookExists == models.Booking{}) {
		var err = fmt.Errorf("booking does not exist")
		return InvoiceRes{}, err
	}

	if err := config.DB.Table("invoices").Create(&req).Error; err != nil {
		return InvoiceRes{}, err
	}

	result := InvoiceRes{
		InvoiceId:  req.ID,
		CustomerId: req.CustomerId,
		EmployeeId: *req.EmployeeId,
		VehicleId:  req.VehicleId,
		TotalPrice: req.TotalPrice,
		Discount:   req.Discount,
		Amount:     req.Amount,
		IsPaid:     req.IsPaid,
		IsCancled:  req.IsCancled,
		EndDate:    req.EndDate,
		BookingId:  req.BookingId,
	}

	return result, nil
}

func (r *invoiceRepository) GetInvoiceByIdRepo(id uint) (InvoiceRes, error) {
	var res models.Invoice

	if err := config.DB.Table("invoices").Where("id = ?", id).Find(&res).Error; err != nil {
		return InvoiceRes{}, err
	}

	if res == (models.Invoice{}) {
		var err = fmt.Errorf("invoice does not exist")
		return InvoiceRes{}, err
	}

	result := InvoiceRes{
		InvoiceId:  res.ID,
		CustomerId: res.CustomerId,
		EmployeeId: *res.EmployeeId,
		VehicleId:  res.VehicleId,
		TotalPrice: res.TotalPrice,
		Discount:   res.Discount,
		Amount:     res.Amount,
		IsPaid:     res.IsPaid,
		IsCancled:  res.IsCancled,
		EndDate:    res.EndDate,
		BookingId:  res.BookingId,
	}

	return result, nil

}

type InvoiceByDateReq struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func (r *invoiceRepository) GetInvoiceByDateRepo(req InvoiceByDateReq) ([]InvoiceRes, error) {
	var res []models.Invoice

	if err := config.DB.Table("invoices").
		Where("end_date BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Find(&res).Error; err != nil {
		return []InvoiceRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("invoice does not exist")
		return []InvoiceRes{}, err
	}

	var result []InvoiceRes
	for _, in := range res {
		tmp := InvoiceRes{
			InvoiceId:  in.ID,
			CustomerId: in.CustomerId,
			EmployeeId: *in.EmployeeId,
			VehicleId:  in.VehicleId,
			TotalPrice: in.TotalPrice,
			Discount:   in.Discount,
			Amount:     in.Amount,
			IsPaid:     in.IsPaid,
			IsCancled:  in.IsCancled,
			EndDate:    in.EndDate,
			BookingId:  in.BookingId,
		}

		result = append(result, tmp)

	}

	return result, nil
}

func (r *invoiceRepository) GetPaidInvoiceByDateRepo(req InvoiceByDateReq) ([]InvoiceRes, error) {
	var res []models.Invoice

	if err := config.DB.Table("invoices").
		Where("end_date BETWEEN ? AND ? AND is_paid = true", req.StartDate, req.EndDate).
		Find(&res).Error; err != nil {
		return []InvoiceRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("invoice does not exist")
		return []InvoiceRes{}, err
	}

	var result []InvoiceRes
	for _, in := range res {
		tmp := InvoiceRes{
			InvoiceId:  in.ID,
			CustomerId: in.CustomerId,
			EmployeeId: *in.EmployeeId,
			VehicleId:  in.VehicleId,
			TotalPrice: in.TotalPrice,
			Discount:   in.Discount,
			Amount:     in.Amount,
			IsPaid:     in.IsPaid,
			IsCancled:  in.IsCancled,
			EndDate:    in.EndDate,
			BookingId:  in.BookingId,
		}

		result = append(result, tmp)

	}

	return result, nil
}

func (r *invoiceRepository) GetCancledInvoiceByDateRepo(req InvoiceByDateReq) ([]InvoiceRes, error) {
	var res []models.Invoice

	if err := config.DB.Table("invoices").
		Where("end_date BETWEEN ? AND ? AND is_cancled = true", req.StartDate, req.EndDate).
		Find(&res).Error; err != nil {
		return []InvoiceRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("invoice does not exist")
		return []InvoiceRes{}, err
	}

	var result []InvoiceRes
	for _, in := range res {
		tmp := InvoiceRes{
			InvoiceId:  in.ID,
			CustomerId: in.CustomerId,
			EmployeeId: *in.EmployeeId,
			VehicleId:  in.VehicleId,
			TotalPrice: in.TotalPrice,
			Discount:   in.Discount,
			Amount:     in.Amount,
			IsPaid:     in.IsPaid,
			IsCancled:  in.IsCancled,
			EndDate:    in.EndDate,
			BookingId:  in.BookingId,
		}

		result = append(result, tmp)

	}

	return result, nil
}

// paid already
func (r *invoiceRepository) GetInvoiceByEmpIdAndDateRepo(id uint, req InvoiceByDateReq) ([]InvoiceRes, error) {
	var res []models.Invoice

	if err := config.DB.Table("invoices").
		Where("employee_id = ? AND is_paid = true AND end_date BETWEEN ? AND ?", id, req.StartDate, req.EndDate).
		Find(&res).Error; err != nil {
		return []InvoiceRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("invoice does not exist")
		return []InvoiceRes{}, err
	}

	var result []InvoiceRes
	for _, in := range res {
		tmp := InvoiceRes{
			InvoiceId:  in.ID,
			CustomerId: in.CustomerId,
			EmployeeId: *in.EmployeeId,
			VehicleId:  in.VehicleId,
			TotalPrice: in.TotalPrice,
			Discount:   in.Discount,
			Amount:     in.Amount,
			IsPaid:     in.IsPaid,
			IsCancled:  in.IsCancled,
			EndDate:    in.EndDate,
			BookingId:  in.BookingId,
		}

		result = append(result, tmp)

	}

	return result, nil

}

func (r *invoiceRepository) GetInvoiceByCusIdRepo(id uint) ([]InvoiceRes, error) {
	var res []models.Invoice

	if err := config.DB.Table("invoices").
		Where("customer_id = ? AND is_paid = true", id).
		Find(&res).Error; err != nil {
		return []InvoiceRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("invoice does not exist")
		return []InvoiceRes{}, err
	}

	var result []InvoiceRes
	for _, in := range res {
		tmp := InvoiceRes{
			InvoiceId:  in.ID,
			CustomerId: in.CustomerId,
			EmployeeId: *in.EmployeeId,
			VehicleId:  in.VehicleId,
			TotalPrice: in.TotalPrice,
			Discount:   in.Discount,
			Amount:     in.Amount,
			IsPaid:     in.IsPaid,
			IsCancled:  in.IsCancled,
			EndDate:    in.EndDate,
			BookingId:  in.BookingId,
		}

		result = append(result, tmp)

	}

	return result, nil
}

// test already
// cus can see their invoice and employee can get to finish pay
func (r *invoiceRepository) GetInvoiceByBookingIdRepo(id uint) (models.Invoice, error) {
	var res models.Invoice

	if err := config.DB.Table("invoices").
		Where("booking_id = ?", id).Find(&res).Error; err != nil {
		return models.Invoice{}, err
	}

	if res == (models.Invoice{}) {
		var err = fmt.Errorf("invoice does not exist")
		return models.Invoice{}, err
	}

	return res, nil
}

// check ispaid false? update enddate and booking isfinished == true
func (r *invoiceRepository) UpdatePaidInvoiceRepo(id uint) (InvoiceRes, error) {
	var res models.Invoice

	if err := config.DB.Table("invoices").
		Where("id = ? AND is_paid = false", id).Find(&res).Error; err != nil {
		return InvoiceRes{}, err
	}

	if res == (models.Invoice{}) {
		var err = fmt.Errorf("invoice does not exist")
		return InvoiceRes{}, err
	}

	if err := config.DB.Table("bookings").
		Where("id = ? AND is_finished = true", res.BookingId).Error; err != nil {
		return InvoiceRes{}, err
	}

	res.IsPaid = true
	res.EndDate = time.Now()

	if err := config.DB.Table("invoices").Save(&res).Error; err != nil {
		return InvoiceRes{}, err
	}

	result := InvoiceRes{
		InvoiceId:  res.ID,
		CustomerId: res.CustomerId,
		EmployeeId: *res.EmployeeId,
		VehicleId:  res.VehicleId,
		TotalPrice: res.TotalPrice,
		Discount:   res.Discount,
		Amount:     res.Amount,
		IsPaid:     res.IsPaid,
		IsCancled:  res.IsCancled,
		EndDate:    res.EndDate,
		BookingId:  res.BookingId,
	}

	return result, nil
}

// enddate
// test already
func (r *invoiceRepository) UpdateCancleInvoiceRepo(bookingId uint) (models.Invoice, error) {
	var res models.Invoice

	if err := config.DB.Table("invoices").
		Where("booking_id = ?", bookingId).Find(&res).Error; err != nil {
		return models.Invoice{}, err
	}

	if res == (models.Invoice{}) {
		var err = fmt.Errorf("invoice does not exist")
		return models.Invoice{}, err
	}

	res.IsCancled = true
	res.EndDate = time.Now()

	if err := config.DB.Table("invoices").Save(&res).Error; err != nil {
		return models.Invoice{}, nil
	}

	return res, nil
}

type UpdateInvoiceReq struct {
	EmployeeId uint
	BookingId  uint
	TotalPrice int `json:"totalPrice"`
	Discount   int `json:"discount"`
}

// test already
func (r *invoiceRepository) UpdateInvoiceRepo(req UpdateInvoiceReq) (models.Invoice, error) {
	var res models.Invoice

	if err := config.DB.Table("invoices").
		Where("booking_id = ? AND employee_id = ? AND is_cancled = false", req.BookingId, req.EmployeeId).
		Find(&res).Error; err != nil {
		return models.Invoice{}, err
	}

	if res == (models.Invoice{}) {
		var err = fmt.Errorf("invoice does not exist")
		return models.Invoice{}, err
	}

	res.TotalPrice = req.TotalPrice
	res.Discount = req.Discount
	res.Amount = req.TotalPrice - req.Discount

	if err := config.DB.Table("invoices").Save(&res).Error; err != nil {
		return models.Invoice{}, err
	}

	return res, nil
}
