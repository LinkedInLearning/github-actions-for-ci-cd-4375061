package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Data struct {
	GUID     string `json:"guid"`
	School   string `json:"school"`
	Mascot   string `json:"mascot"`
	Nickname string `json:"nickname"`
	Location string `json:"location"`
	LatLong  string `json:"latlong"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := setupRouter()

	if err := r.Run(":3000"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", getAllData)
	r.GET("/:guid", getDataByID)

	return r
}

func getAllData(c *gin.Context) {
	var data []Data

	file, err := os.ReadFile("data.json")

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not read data file"})
		return
	}

	err = json.Unmarshal(file, &data)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not unmarshal data"})
		return
	}

	c.JSON(200, data)
}

func getDataByID(c *gin.Context) {
	var data []Data

	guid := c.Param("guid")

	file, err := os.ReadFile("data.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not read data file"})
		return
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not unmarshal data"})
		return
	}

	for _, d := range data {
		if d.GUID == guid {
			c.JSON(200, d)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Data not found"})
}
