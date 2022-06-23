package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// routes "github.com/utsasharma/csv-parsing/routes"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/tvshows", controller.paramControl).Methods("GET")

	// Starting server
	log.Fatal(http.ListenAndServe(":8000", r))

}

// func paramControl(w http.ResponseWriter, r *http.Request) {
// 	no := r.URL.Query().Get("count")
// 	n, _ := strconv.Atoi(no)
// 	// fmt.Println("Value of count is -> ", n)

// 	mType := r.URL.Query().Get("movieType")
// 	// fmt.Println("Value of movie Type is ->", mType)

// 	country := r.URL.Query().Get("country")
// 	// fmt.Println("Value of Country is -> ", country)

// 	startDate := r.URL.Query().Get("startDate")
// 	endDate := r.URL.Query().Get("endDate")
// 	// fmt.Println("Value of startDate is -> ", startDate)
// 	// fmt.Println("Value of endDate is -> ", endDate)

// 	if n != 0 && mType != "" {
// 		getTvshowsByType(w, r, n, mType)
// 	} else if n != 0 && country != "" {
// 		getTvshowsByCountry(w, r, n, country)
// 	} else if n != 0 && startDate != "" && endDate != "" {
// 		getTvshowsByStartDateEndDate(w, r, n, startDate, endDate)
// 	} else if n != 0 {
// 		getTvshows(w, r, n)
// 	}
// }

// // Get 'n' TV shows
// func getTvshows(w http.ResponseWriter, r *http.Request, n int) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Opening csv file
// 	csvfile, err := os.Open("constants/netflix_titles.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// new csv Parsing reader
// 	csvReader := csv.NewReader(csvfile)

// 	var items []models.Item
// 	i := 0
// 	for {
// 		record, err := csvReader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal("File reader error", err)
// 		}

// 		if i < n && record[1] == "TV Show" {
// 			items = append(items, models.Item{
// 				ShowId:       record[0],
// 				Type:         record[1],
// 				Title:        record[2],
// 				Director:     record[3],
// 				Cast:         record[4],
// 				Country:      record[5],
// 				DateAdded:    record[6],
// 				Release_year: record[7],
// 				Rating:       record[8],
// 				Duration:     record[9],
// 				Listed_in:    record[10],
// 				Description:  record[11],
// 			})
// 			i++
// 		}
// 	}
// 	json.NewEncoder(w).Encode(items)
// }

// // Get Tv shows by genre type
// func getTvshowsByType(w http.ResponseWriter, r *http.Request, n int, mType string) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Opening csv file
// 	csvfile, err := os.Open("constants/netflix_titles.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// new csv Parsing reader
// 	csvReader := csv.NewReader(csvfile)

// 	var items []models.Item
// 	i := 0
// 	for {
// 		record, err := csvReader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal("File reader error", err)
// 		}

// 		match, err := regexp.MatchString(mType, record[10])
// 		if match == true && i < n && record[1] == "Movie" {
// 			items = append(items, models.Item{
// 				ShowId:       record[0],
// 				Type:         record[1],
// 				Title:        record[2],
// 				Director:     record[3],
// 				Cast:         record[4],
// 				Country:      record[5],
// 				DateAdded:    record[6],
// 				Release_year: record[7],
// 				Rating:       record[8],
// 				Duration:     record[9],
// 				Listed_in:    record[10],
// 				Description:  record[11],
// 			})
// 			i++
// 		}
// 	}
// 	json.NewEncoder(w).Encode(items)
// }

// // Get tv shows by country origin
// func getTvshowsByCountry(w http.ResponseWriter, r *http.Request, n int, country string) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Opening csv file
// 	csvfile, err := os.Open("constants/netflix_titles.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// new csv Parsing reader
// 	csvReader := csv.NewReader(csvfile)

// 	var items []models.Item
// 	i := 0
// 	for {
// 		record, err := csvReader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if i < n && record[1] == "TV Show" && record[5] == country {
// 			items = append(items, models.Item{
// 				ShowId:       record[0],
// 				Type:         record[1],
// 				Title:        record[2],
// 				Director:     record[3],
// 				Cast:         record[4],
// 				Country:      record[5],
// 				DateAdded:    record[6],
// 				Release_year: record[7],
// 				Rating:       record[8],
// 				Duration:     record[9],
// 				Listed_in:    record[10],
// 				Description:  record[11],
// 			})
// 			i++
// 		}
// 	}
// 	json.NewEncoder(w).Encode(items)
// }

// // get Tv shows within date Range
// func getTvshowsByStartDateEndDate(w http.ResponseWriter, r *http.Request, n int, startDate string, endDate string) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Parsing date taken from params
// 	startDateParsed, _ := time.Parse("2006-02-01", startDate)
// 	endDateParsed, _ := time.Parse("2006-02-01", endDate)

// 	// Changing Parsed date format to match with the csv date extracted
// 	startDateFormatted := startDateParsed.Format("January 2, 2006")
// 	endDateFormatted := endDateParsed.Format("January 2, 2006")

// 	// Changing formatted date type from String to date
// 	startDateFormattedDate, _ := time.Parse("January 2, 2006", startDateFormatted)
// 	endDateFormattedDate, _ := time.Parse("January 2, 2006", endDateFormatted)

// 	// Opening csv file
// 	csvfile, err := os.Open("constants/netflix_titles.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// new csv Parsing reader
// 	csvReader := csv.NewReader(csvfile)

// 	var items []models.Item
// 	i := 0
// 	for {
// 		record, err := csvReader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if i < n && record[1] == "TV Show" {
// 			t, _ := time.Parse("January 2, 2006", record[6])
// 			if t.After(startDateFormattedDate) && t.Before(endDateFormattedDate) {
// 				items = append(items, models.Item{
// 					ShowId:       record[0],
// 					Type:         record[1],
// 					Title:        record[2],
// 					Director:     record[3],
// 					Cast:         record[4],
// 					Country:      record[5],
// 					DateAdded:    record[6],
// 					Release_year: record[7],
// 					Rating:       record[8],
// 					Duration:     record[9],
// 					Listed_in:    record[10],
// 					Description:  record[11],
// 				})
// 				i++
// 			}
// 		}
// 	}
// 	json.NewEncoder(w).Encode(items)
// }
