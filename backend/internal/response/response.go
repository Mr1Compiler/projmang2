package response

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": data})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"success": false, "error": message})
}

func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}

func ValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errs := make(map[string]string)
		for _, e := range ve {
			field := strings.ToLower(e.Field()[:1]) + e.Field()[1:] // camelCase
			errs[field] = formatValidationMessage(e)
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "errors": errs})
		return
	}
	// Handle date/time parsing errors with user-friendly message
	if strings.Contains(err.Error(), "parsing time") {
		BadRequest(c, "invalid date format, use RFC3339 (e.g., 2025-01-15T00:00:00Z)")
		return
	}
	BadRequest(c, err.Error())
}

func formatValidationMessage(e validator.FieldError) string {
	field := strings.ToLower(e.Field()[:1]) + e.Field()[1:]
	switch e.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "min":
		return field + " must be at least " + e.Param() + " characters"
	case "max":
		return field + " must be at most " + e.Param() + " characters"
	default:
		return field + " is invalid"
	}
}
