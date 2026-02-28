package dataLoader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

type FileData []map[string]string

func FetchCSVData(path string) FileData {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	values := FileData{}

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read header
	keys := strings.Split(scanner.Text(), ",")

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		row := map[string]string{}
		for i, value := range data {
			if i < len(keys) {
				row[keys[i]] = value
			}
		}
		values = append(values, row)
	}

	return values
}

// WriteCSV writes FileData to the given path using the provided keys as header order.
// If keys is empty the function will deterministically derive keys from the first row
// and sort them to ensure a stable order.
func WriteCSV(path string, keys []string, data FileData) error {
	if len(data) == 0 {
		return fmt.Errorf("no data to write")
	}

	// determine keys if not provided
	if len(keys) == 0 {
		for k := range data[0] {
			keys = append(keys, k)
		}
		sort.Strings(keys)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// write header
	if err := w.Write(keys); err != nil {
		return err
	}

	// write rows in the order of keys
	for _, row := range data {
		record := make([]string, 0, len(keys))
		for _, k := range keys {
			record = append(record, row[k])
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}
