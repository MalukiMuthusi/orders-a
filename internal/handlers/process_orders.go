package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/orders-a/internal/srv"
	"github.com/gin-gonic/gin"
)

// ProcessOrders start reading and processing orders provided in the csv file
type ProcessOrders struct {
	Srv srv.Srv
}

// Handle start reading and processing orders provided in the csv file
func (h ProcessOrders) Handle(c *gin.Context) {

	// process the orders in background
	go h.Srv.ProcessOrders()

	resp := struct {
		Status string `json:"status"`
	}{
		Status: "PROCESSING",
	}

	c.JSON(http.StatusAccepted, resp)
}
