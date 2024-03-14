package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that creates a new vehicle
	Create(v Vehicle) (err error)
	// FindByColorYear is a method that returns a map of vehicles by color and year
	FindByColorYear(color string, year int) (v map[int]Vehicle, err error)
	// FindByBrandRange is a method that returns a map of vehicles by brand and year range
	FindByBrandRange(brand string, startYear, endYear int) (v map[int]Vehicle, err error)
	//AverageSpeed is a method that returns the average speed of a vehicle by brand
	AverageSpeed(brand string) (average float64, err error)
	// CreateBatch is a method that creates multiple vehicles
	CreateBatch(v []Vehicle) (err error)
	// UpdateSpeed is a method that updates the speed of a vehicle
	UpdateSpeed(id int, speed float64) (err error)
	// FindByFuelType is a method that returns a map of vehicles by fuel type
	FindByFuelType(fuelType string) (v map[int]Vehicle, err error)
	// Delete is a method that deletes a vehicle
	Delete(id int) (err error)
	// UpdateFuelType is a method that updates the fuel type of a vehicle
	UpdateFuelType(id int, fuelType string) (err error)
	// FindByDimensions is a method that returns a map of vehicles by dimensions
	FindByDimensions(minlength, maxlength, minwidth, maxwidth float64) (v map[int]Vehicle, err error)
}
