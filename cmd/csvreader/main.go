package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := flag.String("file", "data/sample.csv", "path to the CSV file")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("unable to open file %s: %v", *filePath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("unable to parse file as CSV: %v", err)
	}

	for i, record := range records {
		fmt.Printf("Record %d: %v\n", i, record)
	}
}
