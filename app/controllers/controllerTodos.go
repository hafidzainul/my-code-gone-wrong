package controllers

import (
	"assesmentbulk/app/models"
	"assesmentbulk/app/syshelper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListTodos function
func (strDB *StrDB) ListTodos(c *gin.Context) {
	var (
		todos  []models.Todos
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	msg := strDB.DB.Find(&todos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&todos)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// CreateTodos function
func (strDB *StrDB) CreateTodos(c *gin.Context) {
	var (
		todos  models.Todos
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	msg := c.Bind(&todos)
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Create(&todos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&todos)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// UpdateTodos function
func (strDB *StrDB) UpdateTodos(c *gin.Context) {
	var (
		todos, newTodos models.Todos
		result          gin.H
		err             syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := c.Bind(&newTodos)
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Find(&todos, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Model(&todos).Updates(newTodos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&todos)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// DeleteTodos function
func (strDB *StrDB) DeleteTodos(c *gin.Context) {
	var (
		todos  models.Todos
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.First(&todos, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	msg = strDB.DB.Delete(&todos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&todos)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// DetailTodos function
func (strDB *StrDB) DetailTodos(c *gin.Context) {
	var (
		todos  models.Todos
		result gin.H
		err    syshelper.ErrorMessage
	)

	defer err.SystemErrorHandler()

	id := c.Param("id")
	msg := strDB.DB.First(&todos, id).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&todos)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}

// SearchTodos function
func (strDB *StrDB) SearchTodos(c *gin.Context) {
	var (
		todos  []models.Todos
		result gin.H
		err    syshelper.ErrorMessage
		search syshelper.SearchQuery
	)

	defer err.SystemErrorHandler()

	search.AppendSearch(c.DefaultQuery("name", ""))
	search.AppendSearch(c.DefaultQuery("description", ""))
	search.AppendSearch(c.DefaultQuery("userid", "0"))

	msg := strDB.DB.Where("name like ? or description like ? or user_id like ?", search.QueryString...).Find(&todos).Error
	if msg != nil {
		err.CustomErrorHandler(msg)
	}

	if len(err.StringError) <= 0 {
		result = syshelper.ResponseSuccess(&todos)
	} else {
		result = syshelper.ResponseFailed(err.StringError)
	}

	c.JSON(http.StatusOK, result)
}
