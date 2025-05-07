package storage

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

// Helper to write a temporary test CSV
func writeTestCSV(filename string, records [][]string) (string, error) {
	path, err := getConfigPath(filename)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return "", err
	}

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return "", err
		}
	}

	return path, nil
}

func TestReadEntriesFromCSV(t *testing.T) {
	records := [][]string{
		{"AcmeCorp", "Fix bug", "45", "2025-04-11"},
		{"Initech", "Build feature", "30", "2025-04-12"},
	}

	_, err := writeTestCSV("entries.csv", records)
	if err != nil {
		t.Fatalf("Failed to write test entries.csv: %v", err)
	}

	entries, err := ReadEntriesFromCSV()
	if err != nil {
		t.Fatalf("ReadEntriesFromCSV returned error: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0].Customer != "AcmeCorp" || entries[0].Minutes != 45 {
		t.Errorf("Unexpected first entry: %+v", entries[0])
	}
}

func TestCustomerFromCSV(t *testing.T) {
	records := [][]string{
		{"AcmeCorp"},
		{"Globex"},
	}

	_, err := writeTestCSV("customers.csv", records)
	if err != nil {
		t.Fatalf("Failed to write test customers.csv: %v", err)
	}

	customers, err := CustomerFromCSV()
	if err != nil {
		t.Fatalf("CustomerFromCSV returned error: %v", err)
	}

	if len(customers) != 2 {
		t.Fatalf("Expected 2 customers, got %d", len(customers))
	}

	if customers[0].Name != "AcmeCorp" {
		t.Errorf("Unexpected customer: %+v", customers[0])
	}
}
