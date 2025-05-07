package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"timekeeper/models"
	"path/filepath"
)

// SaveEntryToCSV saves a time entry to a CSV file
// It creates the file if it doesn't exist and appends to it if it does
// It expects the CSV file to have the following columns: Customer, Task, Minutes, Date
// The Date is expected to be in the format YYYY-MM-DD
func SaveEntryToCSV(entry models.Entry) error {
	path, err := GetConfigPath("entries.csv")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{
		entry.Customer,
		entry.Process,
		entry.Task,
		strconv.Itoa(entry.Minutes),
		entry.Date,
	})
}

// SaveCustomerToCSV saves a customer to a CSV file
// It creates the file if it doesn't exist and appends to it if it does
// It expects the CSV file to have the following columns: Customer
func SaveCustomerToCSV(customer models.Customer) error {
	path, err := GetConfigPath("customers.csv")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{customer.Name})
}

func SaveProcessToCSV(process models.Process) error {
	path, err := GetConfigPath("processes.csv")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{process.Customer, process.Process})
}
