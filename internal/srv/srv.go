package srv

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

type Srv interface {
	Read() ([]byte, error)
}

type Csv struct{}

func (s Csv) Read() error {
	f, err := os.Open("path")
	if err != nil {
		// TODO: log error
		return err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	for {

		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Print(err)
		}

		order := models.Order{}
	}
}
