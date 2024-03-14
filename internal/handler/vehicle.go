package handler

import (
	"app/internal"
	"app/internal/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

var (
	// ErrBadRequest is an error that represents a bad request
	ErrBadRequest = "Datos del vehículo mal formados o incompletos."
	// ErrDuplicatedId is an error that represents a duplicated id
	ErrDuplicatedId = "Identificador del vehículo ya existente."
	//ErrBadCriteria is an error for bad parameters
	ErrBadCriteria = "No se encontraron vehículos con esos criterios."
	//ErrBadSpeed is an error for bad speed
	ErrBadSpeed = "Velocidad mal formada o fuera de rango."
	//ErrBadFuelType is an error for bad fuel type
	ErrBadFuelType = "Tipo de combustible mal formado o no admitido."
	//ErrNotFound is an error for not found
	ErrNotFound = "No se encontró el vehículo."
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

func DataMap(v *map[int]internal.Vehicle) map[int]VehicleJSON {
	data := make(map[int]VehicleJSON)
	for key, value := range *v {
		data[key] = VehicleJSON{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}
	return data
}

func VehicleJSONToVehicle(v VehicleJSON) internal.Vehicle {
	newVehicle := internal.Vehicle{
		Id: v.ID,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           v.Brand,
			Model:           v.Model,
			Registration:    v.Registration,
			Color:           v.Color,
			FabricationYear: v.FabricationYear,
			Capacity:        v.Capacity,
			MaxSpeed:        v.MaxSpeed,
			FuelType:        v.FuelType,
			Transmission:    v.Transmission,
			Weight:          v.Weight,
			Dimensions: internal.Dimensions{
				Height: v.Height,
				Length: v.Length,
				Width:  v.Width,
			},
		},
	}
	return newVehicle
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := DataMap(&v)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		bodyMap := map[string]any{}
		if err = json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		//check if all fields have a value
		err = tools.CheckFieldExistance(bodyMap)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest + " " + err.Error(),
			})
			return
		}

		vehicle := VehicleJSON{}
		err = json.Unmarshal(bytes, &vehicle)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}

		newVehicle := VehicleJSONToVehicle(vehicle)

		err = h.sv.Create(newVehicle)
		if err != nil {
			response.JSON(w, http.StatusConflict, map[string]string{
				"message": ErrDuplicatedId + " " + err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusCreated, map[string]string{
			"message": "Vehículo creado exitosamente.",
		})
	}
}

// GetByColorYear is a method that returns a handler for the route GET /vehicles/color/:color/year/:year
func (h *VehicleDefault) GetByColorYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")
		yearValue, err := strconv.Atoi(year)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadCriteria,
			})
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.FindByColorYear(color, yearValue)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		data := DataMap(&v)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByBrandRange is a method that returns a handler for the route GET /vehicles/brand/:brand/start_year/:start_year/end_year/:end_year
func (h *VehicleDefault) GetByBrandRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		yearStart := chi.URLParam(r, "start_year")
		yearEnd := chi.URLParam(r, "end_year")
		yearStartValue, err := strconv.Atoi(yearStart)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadCriteria,
			})
			return
		}
		yearEndValue, err := strconv.Atoi(yearEnd)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadCriteria,
			})
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.FindByBrandRange(brand, yearStartValue, yearEndValue)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, ErrBadCriteria+" - "+err.Error())
			return
		}

		data := DataMap(&v)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetAverageSpeed is a method that returns a handler for the route GET /vehicles/average-speed/:brand
func (h *VehicleDefault) GetAverageSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")

		// process
		// - get vehicles by color and year
		avg, err := h.sv.AverageSpeed(brand)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		data := fmt.Sprintf("%s - average speed is: %.2f mph", brand, avg)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// CreateBatch is a method that returns a handler for the route POST /vehicles/batch
func (h *VehicleDefault) CreateBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		bodyMap := []map[string]any{}
		if err = json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		// check if all fields have a value
		for _, vehicle := range bodyMap {
			err = tools.CheckFieldExistance(vehicle)
			if err != nil {
				response.JSON(w, http.StatusBadRequest, map[string]string{
					"message": ErrBadRequest + " " + err.Error(),
				})
				return
			}
		}
		vehicles := []VehicleJSON{}
		err = json.Unmarshal(bytes, &vehicles)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		newVehicles := []internal.Vehicle{}
		for _, vehicle := range vehicles {
			newVehicle := VehicleJSONToVehicle(vehicle)
			newVehicles = append(newVehicles, newVehicle)
		}
		err = h.sv.CreateBatch(newVehicles)
		if err != nil {
			response.JSON(w, http.StatusConflict, map[string]string{
				"message": ErrDuplicatedId + " " + err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusCreated, map[string]string{
			"message": "Vehículo creado exitosamente.",
		})
	}
}

// UpdateSpeed is a method that returns a handler for the route PUT /vehicles/:id/speed
func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}

		// read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		bodyMap := map[string]any{}
		if err = json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		newSpeedValue, ok := bodyMap["max_speed"].(float64)
		if !ok || newSpeedValue <= 0 {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadSpeed,
			})
			return
		}

		err = h.sv.UpdateSpeed(id, newSpeedValue)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": ErrNotFound + " " + err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Velocidad del vehículo actualizada exitosamente.",
		})
	}
}

// GetByFuelType is a method that returns a handler for the route GET /vehicles/fuel_type/:fuel_type
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		fuelType := chi.URLParam(r, "type")

		// process
		// - get vehicles by color and year
		v, err := h.sv.FindByFuelType(fuelType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		data := DataMap(&v)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Delete is a method that returns a handler for the route DELETE /vehicles/:id
func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": ErrNotFound + " " + err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Vehículo eliminado exitosamente.",
		})
	}
}

// UpdateFuelType is a method that returns a handler for the route PUT /vehicles/:id/fupdate_fuel
func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}

		// read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		bodyMap := map[string]any{}
		if err = json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		newFuelTypeValue, ok := bodyMap["fuel_type"].(string)
		if !ok || newFuelTypeValue == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadFuelType,
			})
			return
		}

		err = h.sv.UpdateFuelType(id, newFuelTypeValue)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": ErrNotFound + " " + err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Tipo del combustible del vehículo actualizado exitosamente.",
		})
	}
}

// GetByDimensions is a method that returns a handler for the route GET /vehicles/dimensions
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		length := r.URL.Query().Get("length")
		width := r.URL.Query().Get("width")

		lengths := strings.Split(length, "-")
		widths := strings.Split(width, "-")

		minlength, err := strconv.ParseFloat(lengths[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		maxlength, err := strconv.ParseFloat(lengths[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		minwidth, err := strconv.ParseFloat(widths[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}
		maxwidth, err := strconv.ParseFloat(widths[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": ErrBadRequest,
			})
			return
		}

		v, err := h.sv.FindByDimensions(minlength, maxlength, minwidth, maxwidth)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		data := DataMap(&v)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}
