package controllers

import (
	"assesmentbulk/app/models"
	"assesmentbulk/app/syshelper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListOffice function
func (strDB *StrDB) ListOffice(c *gin.Context) {
	var (
		office []models.Office
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	msg := strDB.DB.Find(&office).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&office)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// CreateOffice function
func (strDB *StrDB) CreateOffice(c *gin.Context) {
	var (
		office models.Office
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	msg := c.Bind(&office)
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Create(&office).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&office)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// UpdateOffice function
func (strDB *StrDB) UpdateOffice(c *gin.Context) {
	var (
		office, newOffice models.Office
		result            gin.H
		err               syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := c.Bind(&newOffice)
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.First(&office, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Model(&office).Updates(newOffice).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&office)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// DeleteOffice function
func (strDB *StrDB) DeleteOffice(c *gin.Context) {
	var (
		office models.Office
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.First(&office, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Delete(&office).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(nil)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// DetailOffice function
func (strDB *StrDB) DetailOffice(c *gin.Context) {
	var (
		office []models.Office
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.First(&office, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&office)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// SearchOffice function
func (strDB *StrDB) SearchOffice(c *gin.Context) {
	var (
		office []models.Office
		result gin.H
		err    syshelper.ErrorMessage
		search syshelper.SearchQuery
	)

	defer err.SystemErrorHandler()

	search.AppendSearch(c.DefaultQuery("name", ""))
	search.AppendSearch(c.DefaultQuery("address", ""))

	msg := strDB.DB.Where("name like ? or address like ?", search.QueryString...).Find(&office).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&office)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}
