package vehicle

import (
	"booking-api/config"
	"booking-api/models"
	"fmt"
)

type VehicleRepository interface {
	CreateVehicleRepo(req VehicleReq) (VehicleRes, error)
	GetVehicleByIdRepo(id uint) (VehicleRes, error)
	GetAllVehicleRepo() ([]VehicleRes, error)
	GetVehicleByCusIdRepo(id uint) ([]VehicleRes, error)
	UpdateVehicleRepo(req UpdateVehicleReq) (VehicleRes, error)
	DeleteVehicleRepo(req DeleteVehicleReq) error
}

type vehicleRepository struct{}

func NewVehicleRepository() VehicleRepository {
	return &vehicleRepository{}
}

type VehicleRes struct {
	VehicleId  uint   `json:"vehicleId"`
	Size       string `json:"size"`
	Brand      string `json:"brand"`
	Color      string `json:"color"`
	RegisNo    string `json:"regisNo"`
	CustomerId uint   `json:"customerId"`
}

type VehicleReq struct {
	Size       string `json:"size"`
	Brand      string `json:"brand"`
	Color      string `json:"color"`
	RegisNo    string `json:"regisNo"`
	CustomerId uint
}

func (r *vehicleRepository) CreateVehicleRepo(req VehicleReq) (VehicleRes, error) {
	var userExists models.Customer

	if err := config.DB.Table("customers").
		Where("id = ?", req.CustomerId).Find(&userExists).Error; err != nil {
		return VehicleRes{}, err
	}

	if userExists == (models.Customer{}) {
		var err = fmt.Errorf("user does not exist")
		return VehicleRes{}, err
	}

	newVehicle := models.Vehicle{
		Size:       req.Size,
		Brand:      req.Brand,
		Color:      req.Color,
		RegisNo:    req.RegisNo,
		CustomerId: req.CustomerId,
	}

	if err := config.DB.Table("vehicles").Create(&newVehicle).Error; err != nil {
		return VehicleRes{}, err
	}

	res := VehicleRes{
		VehicleId:  newVehicle.ID,
		Size:       newVehicle.Size,
		Brand:      newVehicle.Brand,
		Color:      newVehicle.Color,
		RegisNo:    newVehicle.RegisNo,
		CustomerId: newVehicle.CustomerId,
	}

	return res, nil
}

func (r *vehicleRepository) GetVehicleByIdRepo(id uint) (VehicleRes, error) {
	var res models.Vehicle

	if err := config.DB.Table("vehicles").
		Where("id = ?", id).Find(&res).Error; err != nil {
		return VehicleRes{}, err
	}

	if res == (models.Vehicle{}) {
		var err = fmt.Errorf("vehicle does not exist")
		return VehicleRes{}, err
	}

	result := VehicleRes{
		VehicleId:  res.ID,
		Size:       res.Size,
		Brand:      res.Brand,
		Color:      res.Color,
		RegisNo:    res.RegisNo,
		CustomerId: res.CustomerId,
	}

	return result, nil
}

// test already
func (r *vehicleRepository) GetVehicleByCusIdRepo(id uint) ([]VehicleRes, error) {
	var res []models.Vehicle

	if err := config.DB.Table("vehicles").Where("customer_id = ?", id).Find(&res).Error; err != nil {
		return []VehicleRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("vehicle does not exist")
		return []VehicleRes{}, err
	}

	var result []VehicleRes
	for _, v := range res {
		tmp := VehicleRes{
			VehicleId:  v.ID,
			Size:       v.Size,
			Brand:      v.Brand,
			Color:      v.Color,
			RegisNo:    v.RegisNo,
			CustomerId: v.CustomerId,
		}
		result = append(result, tmp)
	}

	return result, nil

}

// employee
func (r *vehicleRepository) GetAllVehicleRepo() ([]VehicleRes, error) {
	var res []models.Vehicle

	if err := config.DB.Table("vehicles").Find(&res).Error; err != nil {
		return []VehicleRes{}, err
	}

	if len(res) == 0 {
		var err = fmt.Errorf("vehicle does not exist")
		return []VehicleRes{}, err
	}

	var result []VehicleRes
	for _, v := range res {
		tmp := VehicleRes{
			VehicleId:  v.ID,
			Size:       v.Size,
			Brand:      v.Brand,
			Color:      v.Color,
			RegisNo:    v.RegisNo,
			CustomerId: v.CustomerId,
		}
		result = append(result, tmp)
	}

	return result, nil
}

type UpdateVehicleReq struct {
	VehicleId  uint `json:"vehicleId"`
	CustomerId uint
	Size       string `json:"size"`
	Brand      string `json:"brand"`
	Color      string `json:"color"`
	RegisNo    string `json:"regisNo"`
}

func (r *vehicleRepository) UpdateVehicleRepo(req UpdateVehicleReq) (VehicleRes, error) {
	var vehicle models.Vehicle

	if err := config.DB.Table("vehicles").
		Where("id = ? AND customer_id = ?", req.VehicleId, req.CustomerId).Find(&vehicle).Error; err != nil {
		return VehicleRes{}, err
	}

	if vehicle == (models.Vehicle{}) {
		var err = fmt.Errorf("vehicle does not exist")
		return VehicleRes{}, err
	}

	vehicle.Size = req.Size
	vehicle.Brand = req.Brand
	vehicle.Color = req.Color
	vehicle.RegisNo = req.RegisNo

	if err := config.DB.Table("vehicles").Save(&vehicle).Error; err != nil {
		return VehicleRes{}, err
	}

	result := VehicleRes{
		VehicleId:  vehicle.ID,
		Size:       vehicle.Size,
		Brand:      vehicle.Brand,
		Color:      vehicle.Color,
		RegisNo:    vehicle.RegisNo,
		CustomerId: vehicle.CustomerId,
	}

	return result, nil

}

type DeleteVehicleReq struct {
	CustomerId uint
	VehicleId  []uint `json:"vehicleId"`
}

// test already
func (r *vehicleRepository) DeleteVehicleRepo(req DeleteVehicleReq) error {
	var res models.Vehicle

	for _, v := range req.VehicleId {

		if err := config.DB.Table("vehicles").
			Where("id = ? AND customer_id = ?", v, req.CustomerId).
			Delete(&res).Error; err != nil {
			return err
		}
	}

	return nil

}
