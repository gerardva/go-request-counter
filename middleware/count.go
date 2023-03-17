package middleware

import (
	"fmt"
	"github.com/gerardva/go-counter-api/counter"
	"github.com/gin-gonic/gin"
)

func CountMiddleware(count counter.Counter) gin.HandlerFunc {
	return func(c *gin.Context) {
		count.Increment()
		fmt.Printf("Incremented counter")
	}
}
