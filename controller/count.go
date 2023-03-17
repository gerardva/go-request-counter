package controller

import (
	"fmt"
	"github.com/gerardva/go-counter-api/config"
	"github.com/gerardva/go-counter-api/counter"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CountController struct {
	count counter.Counter
}

func NewCountController(count counter.Counter) *CountController {
	return &CountController{
		count: count,
	}
}

func (h CountController) Check(c *gin.Context) {
	instanceCount := h.count.GetInstanceCount()
	totalCount := h.count.GetTotal()
	instanceName := config.GetConfig().Hostname

	output := fmt.Sprintf("You are talking to instance %s.\nThis is request %d to this instance and request %d to the cluster.", instanceName, instanceCount, totalCount)
	c.String(http.StatusOK, output)
}
