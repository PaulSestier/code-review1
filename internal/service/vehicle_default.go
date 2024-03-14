package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that creates a new vehicle
func (s *VehicleDefault) Create(v internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	return
}

// FindByColorYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) FindByColorYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorYear(color, year)
	return
}

// FindByBrandRange is a method that returns a map of vehicles by brand and year range
func (s *VehicleDefault) FindByBrandRange(brand string, startYear, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandRange(brand, startYear, endYear)
	return
}

// AverageSpeed is a method that returns the average speed of a vehicle by brand
func (s *VehicleDefault) AverageSpeed(brand string) (average float64, err error) {
	average, err = s.rp.AverageSpeed(brand)
	return
}

// CreateBatch is a method that creates multiple vehicles
func (s *VehicleDefault) CreateBatch(v []internal.Vehicle) (err error) {
	err = s.rp.CreateBatch(v)
	return
}

// UpdateSpeed is a method that updates the speed of a vehicle
func (s *VehicleDefault) UpdateSpeed(id int, speed float64) (err error) {
	err = s.rp.UpdateSpeed(id, speed)
	return
}

// FindByFuelType is a method that returns a map of vehicles by fuel type
func (s *VehicleDefault) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByFuelType(fuelType)
	return
}

// Delete is a method that deletes a vehicle
func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}

// UpdateFuelType is a method that updates the fuel type of a vehicle
func (s *VehicleDefault) UpdateFuelType(id int, fuelType string) (err error) {
	err = s.rp.UpdateFuelType(id, fuelType)
	return
}

// FindByDimensions is a method that returns a map of vehicles by dimensions
func (s *VehicleDefault) FindByDimensions(minlength, maxlength, minwidth, maxwidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByDimensions(minlength, maxlength, minwidth, maxwidth)
	return
}
