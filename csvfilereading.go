package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	// Opening file
	csvfile, err := os.Open("data/netflix_titles.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Menu driven options
	var choice int
	fmt.Println("")
	fmt.Println("")
	fmt.Println("************************** Menu **************************")
	fmt.Println("1. List the first n records where type: TV Show")
	fmt.Println("2. List the first n records where listed_in: Horror Movies ")
	fmt.Println("3. List the first n type: Movie where country: India")
	fmt.Println("4. Exit")
	fmt.Println("**********************************************************")
	fmt.Print("Please Enter your choice (1, 2, 3 ..): ")
	fmt.Scanf("%d", &choice)

	// new csv Parsing reader
	r := csv.NewReader(csvfile)

	// Switch Case
	switch choice {
	case 1:
		firstnTvShows(r)
	case 2:
		firstnHorrorMovies(r)
	case 3:
		firstnIndianMovies(r)
	default:
		fmt.Println("Exiting...")
		break
	}

}

func userInput(n int, m1 int, y1 int, m2 int, y2 int) (int, int, int, int, int) {
	fmt.Println("")
	fmt.Print("Enter number of records you want to see: ")
	fmt.Scanf("%d", &n)
	fmt.Print("Enter start date month and year: ")
	fmt.Scanf("%d %d", &m1, &y1)
	if m1 < 1 || m1 > 12 {
		fmt.Println("*Month value should be between 1 to 12*")
		userInput(n, m1, y1, m2, y2)
	} else if y1 > time.Now().Year() || y1 < 1888 {
		fmt.Println("Year can not be greater than current year and less than 1888")
		userInput(n, m1, y1, m2, y2)
	}
	fmt.Print("Enter end date month and year: ")
	fmt.Scanf("%d %d", &m2, &y2)
	if m2 < 1 || m2 > 12 {
		fmt.Println("*Month value should be between 1 to 12*")
		userInput(n, m1, y1, m2, y2)
	} else if y2 > time.Now().Year() || y2 < 1888 {
		fmt.Println("Year can not be greater than current year and less than 1888")
		userInput(n, m1, y1, m2, y2)
	}
	return n, m1, y1, m2, y2
}

func firstnTvShows(r *csv.Reader) {
	var n, m1, y1, m2, y2 int
	n, m1, y1, m2, y2 = userInput(n, m1, y1, m2, y2)

	initialTime := time.Now()

	inputStartDate := time.Date(y1, time.Month(m1), 1, 1, 1, 1, 0, time.Local)
	inputEndDate := time.Date(y2, time.Month(m2), 1, 1, 1, 1, 0, time.Local)

	i := 0
	fmt.Println("")
	fmt.Println("S.No\tTitle\t\t\t\t\tDate")
	fmt.Println("=================================================")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if record[1] == "TV Show" && i < n {
			t, _ := time.Parse("January 2, 2006", record[6])
			if t.After(inputStartDate) && t.Before(inputEndDate) { // (start date)  record  (end date)
				fmt.Printf("%s\t%0.20s\t\t\t\t%s", record[0], record[2], record[6])
				fmt.Println("")
				i++
			}
		}
	}
	fmt.Println("")
	fmt.Printf("Total time taken to execute %s", time.Since(initialTime))
	main()
}

func firstnHorrorMovies(r *csv.Reader) {
	var n, m1, y1, m2, y2 int
	n, m1, y1, m2, y2 = userInput(n, m1, y1, m2, y2)

	inputStartDate := time.Date(y1, time.Month(m1), 1, 1, 1, 1, 0, time.Local)
	inputEndDate := time.Date(y2, time.Month(m2), 1, 1, 1, 1, 0, time.Local)

	initialTime := time.Now()

	i := 0
	fmt.Println("")
	fmt.Println("S.No\tTitle\t\t\t\t\tDate")
	fmt.Println("=================================================")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		match, err := regexp.MatchString("Horror Movies", record[10])
		if match == true && i < n {
			t, _ := time.Parse("January 2, 2006", record[6])
			if t.After(inputStartDate) && t.Before(inputEndDate) { // (start date)  record  (end date)
				fmt.Printf("%s\t%0.20s\t\t\t\t%s", record[0], record[2], record[6])
				fmt.Println("")
				i++
			}
		}
	}
	fmt.Println("")
	fmt.Printf("Total time taken to execute %s", time.Since(initialTime))
	main()
}

func firstnIndianMovies(r *csv.Reader) {
	var n, m1, y1, m2, y2 int
	n, m1, y1, m2, y2 = userInput(n, m1, y1, m2, y2)

	inputStartDate := time.Date(y1, time.Month(m1), 1, 1, 1, 1, 0, time.Local)
	inputEndDate := time.Date(y2, time.Month(m2), 1, 1, 1, 1, 0, time.Local)

	initialTime := time.Now()

	i := 0
	fmt.Println("")
	fmt.Println("S.No\tTitle\t\t\t\t\tDate")
	fmt.Println("=================================================")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[5] == "India" && record[1] == "Movie" && i < n {
			t, _ := time.Parse("January 2, 2006", record[6])
			if t.After(inputStartDate) && t.Before(inputEndDate) { // (start date)  record  (end date)
				fmt.Printf("%s\t%0.20s\t\t\t\t%s", record[0], record[2], record[6])
				fmt.Println("")
				i++
			}
		}
	}
	fmt.Println("")
	fmt.Printf("Total time taken to execute %s", time.Since(initialTime))
	main()
}
