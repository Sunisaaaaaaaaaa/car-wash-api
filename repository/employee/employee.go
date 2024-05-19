package employee

import (
	"booking-api/config"
	"booking-api/models"
	"fmt"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

type EmployeeRepository interface {
	GetEmployeeByIdRepo(id uint) (EmployeeRes, error)
	GetAllEmployeeRepo() ([]EmployeeRes, error)
	GetActivedEmployeeRepo() ([]EmployeeRes, error)
	GetInActivedEmployeeRepo() ([]EmployeeRes, error)
	UpdateEmployeeRepo(req EmployeeReq) (EmployeeRes, error)
	DeleteEmployeeRepo(id uint) (EmployeeRes, error)
}

type employeeRepository struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

type EmployeeRes struct {
	EmployeeId uint   `json:"employeeId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"-"` // Omit from JSON response for security reasons
	Phone      string `json:"phone"`
	Role       int    `json:"role"`
	IsActive   bool   `json:"isActive"`
	Gender     string `json:"gender"`
	Age        string `json:"age"`
	Image      string `json:"image"`
}

func (r *employeeRepository) GetEmployeeByIdRepo(id uint) (EmployeeRes, error) {
	var res models.Employee

	if err := config.DB.Table("employees").Where("id = ?", id).Find(&res).Error; err != nil {
		return EmployeeRes{}, err
	}

	if res == (models.Employee{}) {
		var err = fmt.Errorf("employee does not exist")
		return EmployeeRes{}, err
	}

	result := EmployeeRes{
		EmployeeId: res.ID,
		FirstName:  res.FirstName,
		LastName:   res.LastName,
		Email:      res.Email,
		Password:   res.Password,
		Phone:      res.Phone,
		Role:       res.Role,
		IsActive:   res.IsActive,
		Gender:     res.Gender,
		Age:        res.Age,
		Image:      res.Image,
	}

	return result, nil

}

func (r *employeeRepository) GetAllEmployeeRepo() ([]EmployeeRes, error) {
	var res []models.Employee

	if err := config.DB.Table("employees").Find(&res).Error; err != nil {
		return []EmployeeRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("employee does not exist")
		return []EmployeeRes{}, err
	}

	var result []EmployeeRes
	for _, e := range res {
		tmp := EmployeeRes{
			EmployeeId: e.ID,
			FirstName:  e.FirstName,
			LastName:   e.LastName,
			Email:      e.Email,
			Password:   e.Password,
			Phone:      e.Phone,
			Role:       e.Role,
			IsActive:   e.IsActive,
			Gender:     e.Gender,
			Age:        e.Age,
			Image:      e.Image,
		}

		result = append(result, tmp)
	}

	return result, nil

}

func (r *employeeRepository) GetActivedEmployeeRepo() ([]EmployeeRes, error) {
	var res []models.Employee

	if err := config.DB.Table("employees").Where("is_active = true").Find(&res).Error; err != nil {
		return []EmployeeRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("employee does not exist")
		return []EmployeeRes{}, err
	}

	var result []EmployeeRes
	for _, e := range res {
		tmp := EmployeeRes{
			EmployeeId: e.ID,
			FirstName:  e.FirstName,
			LastName:   e.LastName,
			Email:      e.Email,
			Password:   e.Password,
			Phone:      e.Phone,
			Role:       e.Role,
			IsActive:   e.IsActive,
			Gender:     e.Gender,
			Age:        e.Age,
			Image:      e.Image,
		}

		result = append(result, tmp)
	}

	return result, nil

}
func (r *employeeRepository) GetInActivedEmployeeRepo() ([]EmployeeRes, error) {
	var res []models.Employee
	if err := config.DB.Table("employees").Where("is_active = false").Find(&res).Error; err != nil {
		return []EmployeeRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("employee does not exist")
		return []EmployeeRes{}, err
	}

	var result []EmployeeRes
	for _, e := range res {
		tmp := EmployeeRes{
			EmployeeId: e.ID,
			FirstName:  e.FirstName,
			LastName:   e.LastName,
			Email:      e.Email,
			Password:   e.Password,
			Phone:      e.Phone,
			Role:       e.Role,
			IsActive:   e.IsActive,
			Gender:     e.Gender,
			Age:        e.Age,
			Image:      e.Image,
		}

		result = append(result, tmp)
	}

	return result, nil
}

type EmployeeReq struct {
	EmployeeId uint
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Role       int
	IsActive   bool
	Gender     string `json:"gender"`
	Age        string `json:"age"`
	Image      string `json:"image"`
}

func (r *employeeRepository) UpdateEmployeeRepo(req EmployeeReq) (EmployeeRes, error) {
	var before models.Employee

	if err := config.DB.Table("employees").
		Where("id = ? AND is_active = true", req.EmployeeId).
		Find(&before).Error; err != nil {
		return EmployeeRes{}, err
	}

	if before == (models.Employee{}) {
		var err = fmt.Errorf("employee does not exist")
		return EmployeeRes{}, err
	}

	tmp := reflect.TypeOf(req)
	for i := 0; i < tmp.NumField(); i++ {
		field := tmp.Field(i)
		data := reflect.ValueOf(req).FieldByName(field.Name)
		beforePtr := &before

		if data.Kind() == reflect.String && data.String() != "" {
			// fmt.Println(field.Name)
			if field.Name == "Password" {
				encryptedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
				if err != nil {
					return EmployeeRes{}, err
				}
				reflect.ValueOf(beforePtr).Elem().FieldByName(field.Name).SetString(string(encryptedPw))
				continue
			}

			reflect.ValueOf(beforePtr).Elem().FieldByName(field.Name).SetString(data.String())
		}

	}

	if err := config.DB.Table("employees").
		Save(&before).Error; err != nil {
		return EmployeeRes{}, err
	}

	result := EmployeeRes{
		EmployeeId: before.ID,
		FirstName:  before.FirstName,
		LastName:   before.LastName,
		Email:      before.Email,
		Password:   before.Password,
		Phone:      before.Phone,
		Role:       before.Role,
		IsActive:   before.IsActive,
		Gender:     before.Gender,
		Age:        before.Age,
		Image:      before.Image,
	}

	return result, nil

}

func (r *employeeRepository) DeleteEmployeeRepo(id uint) (EmployeeRes, error) {
	var res models.Employee

	if err := config.DB.Table("employees").
		Where("id = ? AND is_active = true", id).
		Find(&res).Error; err != nil {
		return EmployeeRes{}, err
	}

	if res == (models.Employee{}) {
		var err = fmt.Errorf("employee does not exist")
		return EmployeeRes{}, err
	}

	res.IsActive = false

	if err := config.DB.Table("employees").Save(&res).Error; err != nil {
		return EmployeeRes{}, err
	}

	result := EmployeeRes{
		EmployeeId: res.ID,
		FirstName:  res.FirstName,
		LastName:   res.LastName,
		Email:      res.Email,
		Password:   res.Password,
		Phone:      res.Phone,
		Role:       res.Role,
		IsActive:   res.IsActive,
		Gender:     res.Gender,
		Age:        res.Age,
		Image:      res.Image,
	}

	return result, nil

}
