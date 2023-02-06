package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshitbansal05/omdb-api-server/controllers"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	monitor := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	monitor.SetMetricPath("/metrics")
	monitor.Use(router)

	router.GET("/search", controllers.SearchMovies)
	router.GET("/detail/:id", controllers.GetMovieDetailById)
	return router
}
