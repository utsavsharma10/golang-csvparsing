package models

// item struct (Model)
type Item struct {
	ShowId       string `json:"show_id"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	Director     string `json:"director"`
	Cast         string `json:"cast"`
	Country      string `json:"country"`
	DateAdded    string `json:"date_added"`
	Release_year string `json:"release_year"`
	Rating       string `json:"rating"`
	Duration     string `json:"duration"`
	Listed_in    string `json:"listed_in"`
	Description  string `json:"description"`
}
