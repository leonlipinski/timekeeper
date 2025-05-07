package storage

import (
	"encoding/csv"
	// "fmt"
	"os"
	"strconv"
	"timekeeper/models"
)

// ReadEntriesFromCSV reads entries from a CSV file and returns a slice of Entry structs
// It expects the CSV file to have the following columns: Customer, Task, Minutes, Date
func ReadEntriesFromCSV() ([]models.Entry, error) {
	path, _ := GetConfigPath("entries.csv")
	file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }
    var entries []models.Entry
    for _, record := range records {
        if len(record) < 5 {
            continue // skip malformed lines
        }

        minutes, err := strconv.Atoi(record[3])
        if err != nil {
            continue
        }
        entry := models.Entry{
            Customer: record[0],
			Process:  record[1],
            Task:     record[2],
            Minutes:  minutes,
            Date:     record[4],
        }

        entries = append(entries, entry)
    }

    return entries, nil
}

// ReadDatesFromEntriesCSV reads dates from the entries.csv file
// and returns a slice of strings representing the dates
func ReadDatesFromEntriesCSV() ([]string, error) {
	path, _ := GetConfigPath("entries.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var dates []string
	for _, record := range records {
		if len(record) < 4 {
			continue // skip malformed lines
		}

		dates = append(dates, record[3])
	}

	return dates, nil
}

// CustomerFromCSV reads customers from a CSV file
// and returns a slice of Customer structs
// It expects the CSV file to have the following columns: Name
// The CSV file should be located in the config directory
func CustomerFromCSV() ([]models.Customer, error) {
	path, _ := GetConfigPath("customers.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var customers []models.Customer
	for _, record := range records {
		if len(record) < 1 {
			continue // skip malformed lines
		}

		customer := models.Customer{
			Name: record[0],
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func ProcessFromCSV() ([]models.Process, error) {
	path, _ := GetConfigPath("processes.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var processes []models.Process
	for _, record := range records {
		if len(record) < 2 {
			continue // skip malformed lines
		}

		process := models.Process{
			Customer: record[0],
			Process:  record[1],
		}

		processes = append(processes, process)
	}

	return processes, nil
}
