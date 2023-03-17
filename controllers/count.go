package controllers

import (
	"fmt"
	"github.com/gerardva/go-counter-api/config"
	"github.com/gerardva/go-counter-api/counter"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CountController struct{}

func (h CountController) Check(c *gin.Context) {
	instanceCount := counter.GetInstanceCount()
	totalCount := counter.GetTotal()
	instanceName := config.GetConfig().Hostname

	c.String(http.StatusOK, fmt.Sprintf("You are talking to instance %s.\nThis is request %d to this instance and request %d to the cluster.", instanceName, instanceCount, totalCount))
}
