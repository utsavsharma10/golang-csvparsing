package controller

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/utsasharma/csv-parsing/models"
)

func getTvshowsByStartDateEndDate(w http.ResponseWriter, r *http.Request, n int, startDate string, endDate string) {
	w.Header().Set("Content-Type", "application/json")

	// Parsing date taken from params
	startDateParsed, _ := time.Parse("2006-02-01", startDate)
	endDateParsed, _ := time.Parse("2006-02-01", endDate)

	// Changing Parsed date format to match with the csv date extracted
	startDateFormatted := startDateParsed.Format("January 2, 2006")
	endDateFormatted := endDateParsed.Format("January 2, 2006")

	// Changing formatted date type from String to date
	startDateFormattedDate, _ := time.Parse("January 2, 2006", startDateFormatted)
	endDateFormattedDate, _ := time.Parse("January 2, 2006", endDateFormatted)

	// Opening csv file
	csvfile, err := os.Open("constants/netflix_titles.csv")
	if err != nil {
		log.Fatal(err)
	}

	// new csv Parsing reader
	csvReader := csv.NewReader(csvfile)

	var items []models.Item
	i := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i < n && record[1] == "TV Show" {
			t, _ := time.Parse("January 2, 2006", record[6])
			if t.After(startDateFormattedDate) && t.Before(endDateFormattedDate) {
				items = append(items, models.Item{
					ShowId:       record[0],
					Type:         record[1],
					Title:        record[2],
					Director:     record[3],
					Cast:         record[4],
					Country:      record[5],
					DateAdded:    record[6],
					Release_year: record[7],
					Rating:       record[8],
					Duration:     record[9],
					Listed_in:    record[10],
					Description:  record[11],
				})
				i++
			}
		}
	}
	json.NewEncoder(w).Encode(items)
}
