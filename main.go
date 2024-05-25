package main

import (
	"FakeArtifactory/db"
	"FakeArtifactory/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func returnStatic(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func artifacts(c *gin.Context) {
	var artifacts string
	var err error
	artifacts, err = "qwe", nil

	if err != nil {
		badResponse := "Cant get all artifacts"
		c.JSON(http.StatusBadRequest, badResponse)
	} else {
		c.JSON(http.StatusOK, artifacts)
	}
}

func putArtifact(c *gin.Context) {
	requestPath := strings.TrimPrefix(c.Request.URL.Path, "/artifactory/")
	pathParts := strings.Split(requestPath, "/")
	fileName := pathParts[len(pathParts)-1]

	log.Info("File: " + fileName + " Path: " + requestPath)

	// Check if the storage directory exists
	if _, err := os.Stat("storage"); err != nil {
		err := os.Mkdir("storage", 0755)
		if err != nil {
			log.Error("Failed to create storage directory")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create storage directory"})
			return
		}
	}

	filePath := filepath.Join("storage", fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		log.Error("Failed to create file: " + fileName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, c.Request.Body)

	if err != nil {
		log.Error("Failed to save file: " + fileName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	log.Success("File " + fileName + " saved successfully")
	c.JSON(http.StatusOK, gin.H{"message": "File saved successfully"})
}

func dbInitialization() {
	db.Initialize()
}

func main() {
	dbInitialization()

	r := gin.Default()
	r.Static("/img", "./templates/img")
	r.LoadHTMLGlob("templates/index.html")
	r.GET("/", returnStatic)
	r.GET("/ping", pong)

	r.GET("/artifacts", artifacts)
	r.PUT("/artifactory/*path", putArtifact)

	r.Run()
}
