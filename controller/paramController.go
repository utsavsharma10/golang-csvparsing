package controller

import (
	"net/http"
	"strconv"
)

func paramControl(w http.ResponseWriter, r *http.Request) {
	no := r.URL.Query().Get("count")
	n, _ := strconv.Atoi(no)
	// fmt.Println("Value of count is -> ", n)

	mType := r.URL.Query().Get("movieType")
	// fmt.Println("Value of movie Type is ->", mType)

	country := r.URL.Query().Get("country")
	// fmt.Println("Value of Country is -> ", country)

	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	// fmt.Println("Value of startDate is -> ", startDate)
	// fmt.Println("Value of endDate is -> ", endDate)

	if n != 0 && mType != "" {
		getTvshowsByType(w, r, n, mType)
	} else if n != 0 && country != "" {
		getTvshowsByCountry(w, r, n, country)
	} else if n != 0 && startDate != "" && endDate != "" {
		getTvshowsByStartDateEndDate(w, r, n, startDate, endDate)
	} else if n != 0 {
		getTvshows(w, r, n)
	}
}
