package routes

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := strings.TrimSpace(c.Query("name"))
	if name == "" {
		name = "Skilluv"
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     fmt.Sprintf("Hello %s!", name),
		"server_time": time.Now().UTC().Format(time.RFC3339),
	})
}
