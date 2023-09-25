// main.go

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to debug mode
	gin.SetMode(gin.DebugMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.SetFuncMap(template.FuncMap{
		"jsonMarshallIndent": func(data interface{}) string {
			json, err := json.MarshalIndent(data, "", "\t")
			if err != nil {
				fmt.Println(err)
			}
			return string(json)
		}})
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Serve static files
	router.Static("/static", "./static")

	// Start serving the application
	err := router.RunTLS(":3000", "yotiSelfSignedCert.pem", "yotiSelfSignedKey.pem")
	if err != nil {
		panic(err)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	sessionCreatedInterface, _ := c.Get("session_created")
	data["session_created"] = sessionCreatedInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
