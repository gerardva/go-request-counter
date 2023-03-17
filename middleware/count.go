package middleware

import (
	"fmt"
	"github.com/gerardva/go-counter-api/counter"
	"github.com/gin-gonic/gin"
)

func CountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		counter.Increment()
		fmt.Printf("Incremented counter")
	}
}
