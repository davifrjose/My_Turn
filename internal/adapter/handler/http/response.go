package http

import (
	"errors"
	"net/http"

	"github.com/davifrjose/My_Turn/internal/core/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var errorStatusMap = map[error]int{
	model.ErrorInternal:        http.StatusInternalServerError,
	model.ErrorConflictingData: http.StatusConflict,
}

type errorResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func newErrorResponse(errors []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errors,
	}
}

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func validationError(ctx *gin.Context, err error) {
	errMessages := parseError(err)
	response := newErrorResponse(errMessages)

	ctx.JSON(http.StatusBadRequest, response)
}

func parseError(err error) []string {
	var errorMessages []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, err.Error())
		}
	} else {
		errorMessages = append(errorMessages, err.Error())
	}

	return errorMessages
}

func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errorMessage := parseError(err)
	response := newErrorResponse(errorMessage)
	ctx.JSON(statusCode, response)
}

func handleSuccess(ctx *gin.Context, data any) {
	response := newResponse(true, "Operation completed successfully", data)

	ctx.JSON(http.StatusOK, response)
}
