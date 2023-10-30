package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	_ "github.com/muesli/termenv"
)

func main() {
	r := gin.Default()

	r.GET("/nytimes", func(c *gin.Context) {
		// Fetch the RSS feed
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL("https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error fetching RSS feed",
			})
			return
		}

		// Create a slice to store feed items
		var items []gin.H

		// Iterate through feed items and add them to the slice
		for _, item := range feed.Items {
			items = append(items, gin.H{
				"title":       item.Title,
				"description": item.Description,
				"link":        item.Link,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"title": feed.Title,
			"items": items,
		})
	})

	port := ":8080"
	fmt.Println("Server running on", port)
	r.Run(port)
}
