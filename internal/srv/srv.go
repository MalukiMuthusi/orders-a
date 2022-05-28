package srv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/MalukiMuthusi/orders-a/internal/logger"
	"github.com/MalukiMuthusi/orders-a/internal/models"
	"github.com/MalukiMuthusi/orders-a/internal/utils"
	"github.com/spf13/viper"
)

// Srv is a service to manage the csv records
type Srv interface {

	// ProcessOrders reads csv orders data and unmarshal the data
	ProcessOrders() ([]*models.Order, error)
}

// Csv refers to a csv service that reads and decodes csv data
type Csv struct{}

// ProcessOrders reads csv orders data and unmarshal the data
func (s Csv) ProcessOrders() ([]*models.Order, error) {

	f, err := os.Open(viper.GetString(utils.OrdersPath))
	if err != nil {

		logger.Log.Info(err)

		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	var orders []*models.Order

	countryCodes, err := GetCountryCodes()
	if err != nil {
		return nil, err
	}

	for i := 0; ; i++ {

		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			logger.Log.Info(err)
			continue
		}

		// Skip the csv header
		if i == 0 {
			continue
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {

			logger.Log.Info(err)

			continue
		}

		parcelWeight, err := strconv.ParseFloat(record[3], 32)
		if err != nil {

			logger.Log.Info(err)

			continue
		}

		order := &models.Order{
			ID:           id,
			Email:        record[1],
			PhoneNumber:  record[2],
			ParcelWeight: float32(parcelWeight),
		}

		// Get the country for the order, based on the phone number
		country, err := Country(order.PhoneNumber, countryCodes)
		if err != nil {
			continue
		}

		order.Country = country

		orders = append(orders, order)
	}

	return orders, nil
}

// GetCountryCodes returns a map of a country and its phone number REGEXP.
// The values are provided through a csv file
func GetCountryCodes() (map[string]string, error) {

	f, err := os.Open(viper.GetString(utils.CountryCodesPath))
	if err != nil {

		logger.Log.Info(err)

		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	countryCodes := make(map[string]string)

	for i := 0; ; i++ {

		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			logger.Log.Info(err)
			continue
		}

		// Skip the csv header
		if i == 0 {
			continue
		}

		countryCodes[record[0]] = record[1]
	}

	return countryCodes, nil
}

// Country finds the REGEXP that matches the provided phone number
func Country(phoneNumber string, countryCodes map[string]string) (*string, error) {

	for k, v := range countryCodes {

		matched, err := regexp.MatchString(v, phoneNumber)
		if err != nil {
			return nil, err
		}

		if matched {
			return &k, nil
		}
	}

	return nil, utils.ErrNoMatchForPhoneNumberRE
}
