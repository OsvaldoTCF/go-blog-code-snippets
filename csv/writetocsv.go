package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Create("data2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{
		{"name", "age", "city", "job", "company", "phone"},
		{"Kevin", "30", "New York", "Manager", "Folio Inc.", "355-025-0525"},
		{"John", "25", "Los Angeles", "Developer", "Folio Inc.", "555-230-2139"},
		{"Mary", "28", "Chicago", "HR", "Folio Inc.", "331-359-8311"},
		{"Elizabeth", "35", "New York", "Data Analyst", "Apple Inc.", "113-602-9009"},
	}

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data written to csv file successfully.")
}
