package controllers

import (
	"assesmentbulk/app/models"
	"assesmentbulk/app/syshelper"
	"net/http"
	"os"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// UserOnOffice function
func (strDB *StrDB) UserOnOffice(c *gin.Context) {
	var (
		officeuser []models.OfficeUser
		office     []models.Office
		result     gin.H
		err        syshelper.ErrorMessage
		table      syshelper.ReportTable
		tempDB     *gorm.DB
	)

	defer err.SystemErrorHandler()

	tempDB = strDB.DB.Model(&office).Select("offices.name, offices.address, users.fullname, users.email").Joins("inner join users on users.office_id = offices.id")

	msg := tempDB.Count(&table.Total).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	var errTemps []error
	table.CurrentPage, table.PerPage, errTemps = GetQueryStringParameter(c)
	if len(errTemps) > 0 {
		for _, tempDB := range errTemps {
			err.CustomErrorHandler(tempDB)
		}
	}

	table.Path = os.Getenv("URL_STAGING") + "/custom/user-office"
	table.CalculatePage()
	table.SetURL()

	table.Data = &officeuser

	msg = tempDB.Limit(table.PerPage).Offset(table.From - 1).Scan(&officeuser).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		// result = syshelper.ResponseSuccess(&officeuser)
		result = table.ResponseReport()
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// UserJobs function
func (strDB *StrDB) UserJobs(c *gin.Context) {
	var (
		usertodos []models.UserTodos
		user      []models.User
		result    gin.H
		err       syshelper.ErrorMessage
		msg       error
		table     syshelper.ReportTable
		tempDB    *gorm.DB
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")

	if id == "" {
		tempDB = strDB.DB.Model(&user).Select("users.fullname, users.email, todos.name, todos.description").Joins("inner join todos on todos.user_id = users.id")
	} else {
		tempDB = strDB.DB.Model(&user).Select("users.fullname, users.email, todos.name, todos.description").Joins("inner join todos on todos.user_id = users.id").Where("users.id = ?", id)
	}

	msg = tempDB.Count(&table.Total).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	var errTemps []error
	table.CurrentPage, table.PerPage, errTemps = GetQueryStringParameter(c)
	if len(errTemps) > 0 {
		for _, tempDB := range errTemps {
			err.CustomErrorHandler(tempDB)
		}
	}

	table.Path = os.Getenv("URL_STAGING") + "/custom/user-jobs"
	table.CalculatePage()
	table.SetURL()

	msg = tempDB.Limit(table.PerPage).Offset(table.From - 1).Scan(&usertodos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	table.Data = &usertodos

	if len(err.StringError) <= 0 {
		// result = syshelper.ResponseSuccess(&usertodos)
		result = table.ResponseReport()
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// OfficeJobs function
func (strDB *StrDB) OfficeJobs(c *gin.Context) {
	var (
		officetodos []models.OfficeTodos
		office      []models.Office
		result      gin.H
		err         syshelper.ErrorMessage
		msg         error
		table       syshelper.ReportTable
		tempDB      *gorm.DB
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")

	if id == "" {
		tempDB = strDB.DB.Model(&office).Select("offices.name as name, offices.address as address, todos.name as jobs, todos.description as description").Joins("inner join users on users.office_id = offices.id").Joins("inner join todos on todos.user_id = users.id")
	} else {
		tempDB = strDB.DB.Model(&office).Select("offices.name as name, offices.address as address, todos.name as jobs, todos.description as description").Joins("inner join users on users.office_id = offices.id").Joins("inner join todos on todos.user_id = users.id").Where("offices.id = ?", id)
	}

	msg = tempDB.Count(&table.Total).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	var errTemps []error
	table.CurrentPage, table.PerPage, errTemps = GetQueryStringParameter(c)
	if len(errTemps) > 0 {
		for _, tempDB := range errTemps {
			err.CustomErrorHandler(tempDB)
		}
	}

	table.Path = os.Getenv("URL_STAGING") + "/custom/office-jobs"
	table.CalculatePage()
	table.SetURL()

	msg = tempDB.Scan(&officetodos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	table.Data = &officetodos

	if len(err.StringError) <= 0 {
		// result = syshelper.ResponseSuccess(&officetodos)
		result = table.ResponseReport()
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// OfficeUserJobs function
func (strDB *StrDB) OfficeUserJobs(c *gin.Context) {
	var (
		officetodos []models.OfficeTodos
		office      []models.Office
		result      gin.H
		err         syshelper.ErrorMessage
		msg         error
		// table       syshelper.ReportTable
	)

	defer err.SystemErrorHandler()

	officeid := c.Query("officeid")
	userid := c.Query("userid")

	msg = strDB.DB.Model(&office).Select("offices.name as name, offices.address as address, todos.name as jobs, todos.description as description").Joins("inner join users on users.office_id = offices.id").Joins("inner join todos on todos.user_id = users.id").Where("offices.id = ? or users.id = ?", officeid, userid).Scan(&officetodos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&officetodos)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// OfficeByUser function
func (strDB *StrDB) OfficeByUser(c *gin.Context) {
	var (
		office models.Office
		result gin.H
		err    syshelper.ErrorMessage
		// table  syshelper.ReportTable
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.Joins("join users on users.office_id = offices.id").Where("users.id = ?", id).Find(&office).Error
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

// UserByJob function
func (strDB *StrDB) UserByJob(c *gin.Context) {
	var (
		user   models.User
		result gin.H
		err    syshelper.ErrorMessage
		// table  syshelper.ReportTable
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.Joins("join todos on todos.user_id = users.id").Where("todos.id = ?", id).Find(&user).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&user)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// OfficeByJob function
func (strDB *StrDB) OfficeByJob(c *gin.Context) {
	var (
		office models.Office
		result gin.H
		err    syshelper.ErrorMessage
		// table  syshelper.ReportTable
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.Joins("join users on users.office_id = offices.id").Joins("join todos on todos.user_id = users.id").Where("todos.id = ?", id).Find(&office).Error
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

// GetQueryStringParameter function
func GetQueryStringParameter(c *gin.Context) (int, int, []error) {
	var (
		msg []error
	)

	currentPage, temp1 := strconv.Atoi(c.DefaultQuery("page", "1"))
	if temp1 != nil {
		msg = append(msg, temp1)
	}

	perPage, temp2 := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	if temp1 != nil {
		msg = append(msg, temp2)
	}

	return currentPage, perPage, msg
}
