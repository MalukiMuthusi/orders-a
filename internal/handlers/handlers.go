package handlers

import "github.com/gin-gonic/gin"

// DebugPrintRouteFunc prints the available API endpoints
type DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

// SetUpRouter configures the API server
func SetUpRouter(debugPrintRoute DebugPrintRouteFunc) *gin.Engine {

	r := gin.New()

	// log the endpoints
	gin.DebugPrintRouteFunc = debugPrintRoute

	// process orders

	process := ProcessOrders{}

	r.GET("process", process.Handle)

	return r
}
