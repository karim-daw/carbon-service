package helpers

import (
	"github.com/gin-gonic/gin"
)

// respondWithError standardizes error responses.
func RespondWithError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{"error": message})
}
