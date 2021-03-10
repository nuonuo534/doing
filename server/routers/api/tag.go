package api

import (
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/setting"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nuonuo534/doing/pkg/e"
	"github.com/nuonuo534/doing/pkg/util"
	"github.com/unknwon/com"
)

var validate *validator.Validate

// GetTags list
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// AddTag create
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	validate = validator.New()
	tag := &models.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}

	err := validate.Struct(tag)
	code := e.INVALID_PARAMS
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			log.Fatalf(err.Namespace())
			log.Fatalf(err.Field())
			log.Fatalf(err.StructNamespace())
			log.Fatalf(err.StructField())
			log.Fatalf(err.Tag())
			log.Fatalf(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		code = e.ERROR_PARAMS
		// from here you can create your own error messages in whatever language you wish
		// return
	} else {

		if !models.ExistTagByName(name) {
			models.AddTag(name, state, createdBy)
			code = e.SUCCESS

		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag edit
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	createdBy := c.Query("created_by")
	state := c.Query("state")
	var tags models.Tag
	if tags.ExistTableByKey("id", id) {
		data := make(map[string]interface{})
		data["name"] = name
		data["created_by"] = createdBy
		data["state"] = state
		// models.EditTag(id, data)
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTag delete
func DeleteTag(c *gin.Context) {

}
