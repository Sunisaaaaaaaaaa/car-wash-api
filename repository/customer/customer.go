package customer

import (
	"booking-api/config"
	"booking-api/models"
	"fmt"
	"reflect"

	"golang.org/x/crypto/bcrypt"
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

	if res == (models.Customer{}) {
		var err = fmt.Errorf("user does not exists")
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

	if len(res) == 0 {
		var err = fmt.Errorf("user does not exists")
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
	var before models.Customer

	if err := config.DB.Table("customers").
		Where("id = ?", req.ID).
		Find(&before).Error; err != nil {
		return CustomerRes{}, err
	}

	if before == (models.Customer{}) {
		var err = fmt.Errorf("user does not exists")
		return CustomerRes{}, err
	}

	tmp := reflect.TypeOf(req)
	for i := 0; i < tmp.NumField(); i++ {
		field := tmp.Field(i)
		data := reflect.ValueOf(req).FieldByName(field.Name)
		beforePtr := &before

		if data.Kind() == reflect.String && data.String() != "" {
			fmt.Println(field.Name)
			if field.Name == "Password" {
				encryptedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
				if err != nil {
					return CustomerRes{}, err
				}
				reflect.ValueOf(beforePtr).Elem().FieldByName(field.Name).SetString(string(encryptedPw))
				continue
			}

			reflect.ValueOf(beforePtr).Elem().FieldByName(field.Name).SetString(data.String())
		}

	}

	if err := config.DB.Table("customers").
		Save(&before).Error; err != nil {
		config.DB.Rollback()
		return CustomerRes{}, err
	}

	result := CustomerRes{
		ID:        before.ID,
		FirstName: before.FirstName,
		LastName:  before.LastName,
		Email:     before.Email,
		Phone:     before.Phone,
		Role:      before.Role,
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
