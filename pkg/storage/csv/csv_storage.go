/*
		csv 
This package is to be used only by the init file of the Garage 
 */
package csv

import (
	"os"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// CsvStorage ...
type CsvStorage struct {
	CsvFile *os.File
}

func NewCsvStorage() *CsvStorage {
	storage := &CsvStorage{}
	var err error
	storage.CsvFile, err = os.Open("garage.csv")
	if err != nil {
		return nil
	}
	return storage
}

// RegisterGarage ...
func (csv *CsvStorage) RegisterGarage(garage *model.Garage) error {
	return nil
}
