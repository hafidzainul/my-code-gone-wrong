package syshelper

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess function
func ResponseSuccess(data interface{}) gin.H {
	return gin.H{
		"status": "Success",
		"data":   data,
	}
}

// ResponseFailed function
func ResponseFailed(strError []error) gin.H {
	var errs []string

	for _, str := range strError {
		errs = append(errs, str.Error())
	}

	return gin.H{
		"status":  "Errors",
		"message": errs,
	}
}

// ReportTable strunct
type ReportTable struct {
	Total        int64 //
	PerPage      int   //
	CurrentPage  int   //
	LastPage     int
	FirstPageURL string
	LastPageURL  string
	NextPageURL  string
	PrevPageURL  string
	Path         string //
	From         int
	To           int
	Data         interface{} //
}

// CalculatePage function
func (rT *ReportTable) CalculatePage() {
	rT.LastPage = int(math.Ceil(float64(rT.Total) / float64(rT.PerPage)))
	rT.From = (rT.CurrentPage-1)*rT.PerPage + 1
	rT.To = rT.From + rT.PerPage - 1
}

// SetURL function
func (rT *ReportTable) SetURL() {
	var perPage string

	if rT.PerPage != 5 {
		perPage = "&per_page=" + fmt.Sprint(rT.PerPage)
	} else {
		perPage = ""
	}

	rT.LastPageURL = rT.Path + "?page=" + fmt.Sprint(rT.LastPage) + perPage
	rT.FirstPageURL = rT.Path + "?page=" + fmt.Sprint(1) + perPage

	if rT.CurrentPage == rT.LastPage {
		rT.NextPageURL = ""
		rT.PrevPageURL = rT.Path + "?page=" + fmt.Sprint(rT.CurrentPage-1) + perPage
	} else if rT.CurrentPage == 1 {
		rT.NextPageURL = rT.Path + "?page=" + fmt.Sprint(rT.CurrentPage+1) + perPage
		rT.PrevPageURL = ""
	} else {
		rT.NextPageURL = rT.Path + "?page=" + fmt.Sprint(rT.CurrentPage+1) + perPage
		rT.PrevPageURL = rT.Path + "?page=" + fmt.Sprint(rT.CurrentPage-1) + perPage
	}

}

// ResponseReport function
func (rT *ReportTable) ResponseReport() gin.H {
	return gin.H{
		"total":          rT.Total,
		"per_page":       rT.PerPage,
		"current_page":   rT.CurrentPage,
		"last_page":      rT.LastPage,
		"first_page_url": rT.FirstPageURL,
		"last_page_url":  rT.LastPageURL,
		"next_page_url":  rT.NextPageURL,
		"prev_page_url":  rT.PrevPageURL,
		"path":           rT.Path,
		"from":           rT.From,
		"to":             rT.To,
		"data":           rT.Data,
	}
}
