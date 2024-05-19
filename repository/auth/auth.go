package auth

import (
	"booking-api/config"
	"booking-api/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthenRepository interface {
	RegisterRepository(user CustomerRegRequest) (models.Customer, error)
	LoginRepository(user LoginRequest) (string, error)
	RegisterForEmployeeRepository(user EmployeeRegRequest) (models.Employee, error)
	LoginForEmployeeRepository(user LoginRequest) (string, error)
}

type authenRepository struct{}

func NewCustomerRepository() AuthenRepository {
	return &authenRepository{}
}

type CustomerRegRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LsstName  string `json:"lastname" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

func (r *authenRepository) RegisterRepository(user CustomerRegRequest) (models.Customer, error) {
	var userExists models.Customer

	if err := config.DB.Table("customers").
		Where("email = ?", user.Email).
		Find(&userExists).Error; err != nil {
		return models.Customer{}, err
	}

	if userExists != (models.Customer{}) {
		var err = fmt.Errorf("user exists")
		return userExists, err
	} else {
		encryptedPw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		newUser := models.Customer{
			FirstName: user.FirstName,
			LastName:  user.LsstName,
			Email:     user.Email,
			Password:  string(encryptedPw),
			Phone:     user.Phone,
			Role:      2, //customer
		}
		if err := config.DB.Table("customers").Create(&newUser).Error; err != nil {
			return models.Customer{}, err
		} else {
			return newUser, nil
		}
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *authenRepository) LoginRepository(user LoginRequest) (string, error) {
	var userExists models.Customer

	if err := config.DB.Table("customers").
		Where("email = ?", user.Email).
		Find(&userExists).Error; err != nil {
		return "", err
	}

	if userExists == (models.Customer{}) {
		var err = fmt.Errorf("user does not exist")
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(user.Password))
	if err == nil {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExists.ID,
			"role":   userExists.Role,
			"exp":    time.Now().Add(time.Hour * 10).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		if err != nil {
			return "", fmt.Errorf("error to sign token")
		}
		// fmt.Println(tokenString, err)

		return tokenString, nil
	} else {
		return "", fmt.Errorf("loggin failed")
	}
}

type EmployeeRegRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Age       string `json:"age" binding:"required"`
	Image     string `json:"image" binding:"required"`
}

func (r *authenRepository) RegisterForEmployeeRepository(user EmployeeRegRequest) (models.Employee, error) {
	var userExists models.Employee

	if err := config.DB.Table("employees").
		Where("email = ?", user.Email).
		Find(&userExists).Error; err != nil {
		return models.Employee{}, err
	}

	if userExists != (models.Employee{}) {
		var err = fmt.Errorf("user exists")
		return userExists, err
	} else {
		encryptedPw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		newUser := models.Employee{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  string(encryptedPw),
			Phone:     user.Phone,
			Role:      1,
			IsActive:  true,
			Gender:    user.Gender,
			Age:       user.Age,
			Image:     user.Image,
		}
		if err := config.DB.Table("employees").Create(&newUser).Error; err != nil {
			return models.Employee{}, err
		} else {
			return newUser, nil
		}
	}
}

func (r *authenRepository) LoginForEmployeeRepository(user LoginRequest) (string, error) {
	var userExists models.Employee

	if err := config.DB.Table("employees").
		Where("email = ? AND is_active = true", user.Email).
		Find(&userExists).Error; err != nil {
		return "", err
	}

	if userExists == (models.Employee{}) {
		var err = fmt.Errorf("user does not exist")
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(user.Password))
	if err == nil {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExists.ID,
			"role":   userExists.Role,
			"exp":    time.Now().Add(time.Minute * 100).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		if err != nil {
			return "", fmt.Errorf("error to sign token")
		}
		// fmt.Println(tokenString, err)

		return tokenString, nil
	} else {
		return "", fmt.Errorf("loggin failed")
	}
}
