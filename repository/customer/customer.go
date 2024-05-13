package customer

import (
	"booking-api/config"
	"booking-api/models"
)

type CustomerRepository interface {
	GetCustomerByIdRepository(id uint) (CustomerRes, error)
	GetAllCustomerRepository() ([]CustomerRes, error)
	UpdateCustomerRepository(req CustomerUpdateReq) (CustomerRes, error)
	DeleteCustomerRepository(id uint) error
}

type customerRepository struct{}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}

type CustomerRes struct {
	ID        uint   `json:"customerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      int    `json:"role"`
}

func (c *customerRepository) GetCustomerByIdRepository(id uint) (CustomerRes, error) {
	var res models.Customer

	if err := config.DB.Table("customers").
		Where("id = ?", id).
		Find(&res).Error; err != nil {
		return CustomerRes{}, err
	}

	result := CustomerRes{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Phone:     res.Phone,
		Role:      res.Role,
	}

	return result, nil
}

func (c *customerRepository) GetAllCustomerRepository() ([]CustomerRes, error) {
	var res []models.Customer

	if err := config.DB.Table("customers").
		Find(&res).Error; err != nil {
		return []CustomerRes{}, err
	}

	var result []CustomerRes
	for _, cus := range res {
		result = append(result,
			CustomerRes{
				ID:        cus.ID,
				FirstName: cus.FirstName,
				LastName:  cus.LastName,
				Email:     cus.Email,
				Phone:     cus.Phone,
				Role:      cus.Role,
			})
	}

	return result, nil
}

type CustomerUpdateReq struct {
	ID        uint   `json:"customerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Password  string `json:"Password"`
}

func (c *customerRepository) UpdateCustomerRepository(req CustomerUpdateReq) (CustomerRes, error) {
	var res models.Customer

	before, err := c.GetCustomerByIdRepository(req.ID)
	if err != nil && before == (CustomerRes{}) {
		return CustomerRes{}, err
	}

	res.FirstName = req.LastName
	res.LastName = req.LastName
	res.Phone = req.Phone
	res.Password = req.Password

	if err := config.DB.Table("customer").
		Save(&res).Error; err != nil {
		config.DB.Rollback()
		return CustomerRes{}, err
	}

	result := CustomerRes{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Phone:     res.Phone,
		Role:      res.Role,
	}

	return result, nil
}

func (c *customerRepository) DeleteCustomerRepository(id uint) error {

	if err := config.DB.Table("customers").
		Delete(&models.Customer{}, id).Error; err != nil {
		return err
	}

	return nil

}
