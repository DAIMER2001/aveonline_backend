package controller

import (
	"encoding/json"
	"fmt"
	http "net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
}

type DatesBetween struct {
	DateInit *time.Time
	DateEnd  *time.Time
}

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func New() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("form")
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate
}

func validateParams(c *gin.Context, params interface{}) error {

	validate := validator.New()
	err := validate.Struct(params)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		if resp := ToErrResponse(err); resp == nil {
			SendError(c, err)
		} else {
			_, err := json.Marshal(&resp)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":  err,
				"status": http.StatusUnprocessableEntity,
				"msg":    "Error en el formato json enviado",
				"resp":   &resp,
			})
		}
		return err
	}
	return nil
}

func ToErrResponse(err error) *ErrResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}
		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s must be a maximum of %s in length", err.Field(), err.Param())
			case "min":
				resp.Errors[i] = fmt.Sprintf("%s must be a minimum of %s in length", err.Field(), err.Param())
			case "url":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid URL", err.Field())
			case "numeric":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid numeric", err.Field())
			case "date":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid date: %s", err.Field(), err.Param())
			default:
				resp.Errors[i] = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}
		return &resp
	}
	return nil
}

func SendResponse(c *gin.Context, status int, data []byte) {
	c.JSON(status, &data)
}

func SendError(c *gin.Context, err error, msg ...string) {
	c.JSON(500, gin.H{
		"error":  err.Error(),
		"status": 500,
		"msg":    msg,
	})
}
