package controllers

import (
	"assesmentbulk/app/models"
	"assesmentbulk/app/syshelper"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func makeRedis(conn redis.Conn, result *gin.H, sec *map[string]interface{}) {

	var cacheredis []byte
	// map to json
	cacheredis, _ = json.Marshal(result)
	// fmt.Println(string(cacheredis))

	// send json to redis
	_, _ = conn.Do("SET", "result", string(cacheredis))

	// get json from redis
	reply, msg := redis.Bytes(conn.Do("GET", "result"))
	if msg != nil {
		fmt.Println(msg)
	}

	// make map definition with interface
	msg = json.Unmarshal(reply, &sec)
	if msg != nil {
		fmt.Println(msg)
	}

	// for key, value := range sec {
	// 	fmt.Println(key, value)
	// }

	_, _ = conn.Do("EXPIRE", "result", "10")
}

// UserOnOfficeWithRedis function
func (strDB *StrDB) UserOnOfficeWithRedis(c *gin.Context) {
	var (
		officeuser []models.OfficeUser
		office     []models.Office
		result     gin.H
		err        syshelper.ErrorMessage
		table      syshelper.ReportTable
		tempDB     *gorm.DB
		sec        map[string]interface{}
	)

	// defer err.SystemErrorHandler()

	pool := redis.NewPool(
		func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		10,
	)
	pool.MaxActive = 0

	conn := pool.Get()
	defer conn.Close()

	reply, _ := redis.Bytes(conn.Do("GET", "result"))
	if reply == nil {
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

		table.Path = os.Getenv("URL_STAGING") + "/custom/user-office-redis"
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

		fmt.Println("Hello")

		makeRedis(conn, &result, &sec)
	} else {
		msg := json.Unmarshal(reply, &sec)
		if msg != nil {
			fmt.Println(msg)
		}
	}

	// c.JSON(http.StatusOK, result)
	c.JSON(http.StatusOK, sec)
}

// UserJobsWithRedis function
func (strDB *StrDB) UserJobsWithRedis(c *gin.Context) {
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

// OfficeJobsWithRedis function
func (strDB *StrDB) OfficeJobsWithRedis(c *gin.Context) {
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

// OfficeUserJobsWithRedis function
func (strDB *StrDB) OfficeUserJobsWithRedis(c *gin.Context) {
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

// OfficeByUserWithRedis function
func (strDB *StrDB) OfficeByUserWithRedis(c *gin.Context) {
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

// UserByJobWithRedis function
func (strDB *StrDB) UserByJobWithRedis(c *gin.Context) {
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

// OfficeByJobWithRedis function
func (strDB *StrDB) OfficeByJobWithRedis(c *gin.Context) {
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
