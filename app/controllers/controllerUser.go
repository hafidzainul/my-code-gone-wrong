package controllers

import (
	"assesmentbulk/app/models"
	"assesmentbulk/app/syshelper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ListUser function
func (strDB *StrDB) ListUser(c *gin.Context) {
	var (
		user   []models.User
		result gin.H
		err    syshelper.ErrorMessage
		table  syshelper.ReportTable
	)

	defer err.SystemErrorHandler()

	msg := strDB.DB.Model(&user).Count(&table.Total).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	var temps []error
	table.CurrentPage, table.PerPage, temps = GetQueryStringParameter(c)
	if len(temps) > 0 {
		for _, temp := range temps {
			err.CustomErrorHandler(temp)
		}
	}

	table.Path = os.Getenv("URL_STAGING") + "/user"
	table.CalculatePage()
	table.SetURL()

	msg = strDB.DB.Limit(table.PerPage).Offset(table.From - 1).Find(&user).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	table.Data = user

	if len(err.StringError) <= 0 {
		// result = syshelper.ResponseSuccess(&user)
		result = table.ResponseReport()
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// CreateUser function
func (strDB *StrDB) CreateUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	msg := c.Bind(&user)
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	hash, msg := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if msg != nil {
		err.CustomErrorHandler(msg)
	}
	user.Password = string(hash)

	msg = strDB.DB.Create(&user).Error
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

// UpdateUser function
func (strDB *StrDB) UpdateUser(c *gin.Context) {
	var (
		user, newUser models.User
		result        gin.H
		err           syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := c.Bind(&newUser)
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Find(&user, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Model(&user).Updates(newUser).Error
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

// DeleteUser function
func (strDB *StrDB) DeleteUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.First(&user, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Delete(&user).Error
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

// DetailUser function
func (strDB *StrDB) DetailUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.First(&user, id).Error
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

// SearchUser function
func (strDB *StrDB) SearchUser(c *gin.Context) {
	var (
		user   []models.User
		result gin.H
		err    syshelper.ErrorMessage
		search syshelper.SearchQuery
	)

	defer err.SystemErrorHandler()

	search.AppendSearch(c.DefaultQuery("fullname", ""))
	search.AppendSearch(c.DefaultQuery("email", ""))
	search.AppendSearch(c.DefaultQuery("officeid", "0"))

	msg := strDB.DB.Where("fullname like ? or email like ? or office_id like ?", search.QueryString...).Find(&user).Error
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
