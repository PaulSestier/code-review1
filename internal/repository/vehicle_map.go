package repository

import (
	"app/internal"
	"fmt"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates a new vehicle
func (r *VehicleMap) Create(v internal.Vehicle) (err error) {
	//check if id already exists
	if _, ok := r.db[v.Id]; ok {
		err = fmt.Errorf("vehicle with id %d already exists", v.Id)
		return
	}
	r.db[v.Id] = v
	return
}

// FindByColorYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) FindByColorYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = fmt.Errorf("no vehicles found with color %s and year %d", color, year)
		return
	}

	return
}

// FindByBrandRange is a method that returns a map of vehicles by brand and year range
func (r *VehicleMap) FindByBrandRange(brand string, yearstart int, yearend int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= yearstart && value.FabricationYear <= yearend {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = fmt.Errorf("brand: %s - year range: %d - %d", brand, yearstart, yearend)
		return
	}

	return
}

// AverageSpeed is a method that returns the average speed of a vehicle by brand
func (r *VehicleMap) AverageSpeed(brand string) (average float64, err error) {
	var totalSpeed float64
	var totalVehicles float64

	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			totalSpeed += value.MaxSpeed
			totalVehicles++
		}
	}

	// check if map is empty
	if totalVehicles == 0 {
		err = fmt.Errorf("no se encontraron vehículos de la marca %s", brand)
		return
	}

	average = totalSpeed / totalVehicles
	return
}

// CreateBatch is a method that creates multiple vehicles
func (r *VehicleMap) CreateBatch(v []internal.Vehicle) (err error) {
	for _, vehicle := range v {
		if _, ok := r.db[vehicle.Id]; ok {
			err = fmt.Errorf("vehicle with id %d already exists", vehicle.Id)
			return
		}
	}
	for _, vehicle := range v {
		err = r.Create(vehicle)
		if err != nil {
			err = fmt.Errorf("error creating vehicle %s - %s with id %d: %w", vehicle.Brand, vehicle.Model, vehicle.Id, err)
			return
		}
	}
	return
}

// UpdateSpeed is a method that updates a vehicle
func (r *VehicleMap) UpdateSpeed(id int, speed float64) (err error) {
	// check if vehicle exists
	vehicle, ok := r.db[id]
	if !ok {
		err = fmt.Errorf("vehicle with id %d not found", id)
		return
	}

	vehicle.MaxSpeed = speed
	r.db[id] = vehicle
	return
}

// FindByFuelType is a method that returns a map of vehicles by fuel type
func (r *VehicleMap) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.FuelType == fuelType {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = fmt.Errorf("no se encontraron vehículos con combustible: %s", fuelType)
		return
	}

	return
}

// Delete is a method that deletes a vehicle
func (r *VehicleMap) Delete(id int) (err error) {
	// check if vehicle exists
	if _, ok := r.db[id]; !ok {
		err = fmt.Errorf("vehicle with id %d not found", id)
		return
	}
	delete(r.db, id)
	return
}

// UpdateFuelType is a method that updates a vehicle
func (r *VehicleMap) UpdateFuelType(id int, fuelType string) (err error) {
	// check if vehicle exists
	vehicle, ok := r.db[id]
	if !ok {
		err = fmt.Errorf("vehicle with id %d not found", id)
		return
	}

	vehicle.FuelType = fuelType
	r.db[id] = vehicle
	return
}

// FindByDimensions is a method that returns a map of vehicles by dimensions
func (r *VehicleMap) FindByDimensions(minlength, maxlength, minwidth, maxwidth float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Length >= minlength && value.Length <= maxlength && value.Width >= minwidth && value.Width <= maxwidth {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = fmt.Errorf("no se encontraron vehículos con esas dimensiones")
		return
	}

	return
}
