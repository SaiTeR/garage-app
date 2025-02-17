package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func Validate(context *gin.Context, request interface{}) {

	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid format or structure",
		})
		fmt.Printf(err.Error())
		context.Abort()
		return
	}

	// Валидация структуры
	if err := validate.Struct(request); err != nil {
		// Форматируем ошибки валидации для удобства
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"validation_errors": errors,
		})

		context.Abort()
		return
	}
}
