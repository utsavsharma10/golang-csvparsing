package controller

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/utsasharma/csv-parsing/models"
)

// Get Tv shows by genre type
func getTvshowsByType(w http.ResponseWriter, r *http.Request, n int, mType string) {
	w.Header().Set("Content-Type", "application/json")

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
			log.Fatal("File reader error", err)
		}

		match, err := regexp.MatchString(mType, record[10])
		if match == true && i < n && record[1] == "Movie" {
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
	json.NewEncoder(w).Encode(items)
}
