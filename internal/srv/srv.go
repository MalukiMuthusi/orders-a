package srv

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"

	"github.com/MalukiMuthusi/orders-a/internal/models"
)

// Srv is a service to manage the csv records
type Srv interface {
	Read() ([]*models.Order, error)
}

// Csv refers to a csv service that reads and decodes csv data
type Csv struct{}

// Read csv data and unmarshal the data
func (s Csv) Read() ([]*models.Order, error) {
	f, err := os.Open("path")
	if err != nil {
		// TODO: log error
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	var orders []*models.Order

	for {

		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Print(err)
		}

		order := &models.Order{
			ID:           record[0],
			Email:        record[1],
			PhoneNumber:  record[2],
			ParcelWeight: record[3],
		}

		// TODO: determine the country from the phone number using the provided REGEX

		orders = append(orders, order)
	}

	return orders, nil
}
