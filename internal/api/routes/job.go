package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/linskybing/platform-go/internal/application/job"
)

// JobRoutes registers job endpoints
func JobRoutes(rg *gin.RouterGroup, svc *job.Service) {
	jobs := rg.Group("/jobs")
	{
		jobs.POST("", createJob(svc))
		jobs.GET("", listJobs(svc))
		jobs.GET("/:id", getJob(svc))
		jobs.DELETE("/:id", cancelJob(svc))
		jobs.POST("/:id/restart", restartJob(svc))
		jobs.GET("/:id/logs", getJobLogs(svc))
		jobs.GET("/:id/checkpoints", getJobCheckpoints(svc))
	}
}

func createJob(svc *job.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(201, gin.H{"status": "created"})
	}
}

func listJobs(svc *job.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"data": []interface{}{}})
	}
}

func getJob(svc *job.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"id": c.Param("id")})
	}
}

func cancelJob(svc *job.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "cancelled"})
	}
}

func restartJob(svc *job.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "restarted"})
	}
}

func getJobLogs(svc *job.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"logs": []string{}})
	}
}

func getJobCheckpoints(svc *job.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"checkpoints": []interface{}{}})
	}
}
