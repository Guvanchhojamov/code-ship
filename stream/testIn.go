package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func startStream(youtubeID string) error {
	cmd := exec.Command("/bin/bash", "stream_test.sh", youtubeID)
	return cmd.Start() // Run in background
}

func main() {
	r := gin.Default()

	r.GET("/stream/:id", func(c *gin.Context) {
		youtubeID := c.Param("id")
		err := startStream(youtubeID)
		if err != nil {
			log.Println("Error starting stream:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start stream"})
			return
		}

		hlsURL := fmt.Sprintf("http://localhost:8080/hls/%s/index.m3u8", youtubeID)
		c.JSON(http.StatusOK, gin.H{"stream_url": hlsURL})
	})

	r.Run(":8081")
}
