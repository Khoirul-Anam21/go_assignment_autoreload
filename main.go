package main

import (
	"encoding/json"
	"go_assignment_jwt/helpers"
	"go_assignment_jwt/models"
	"go_assignment_jwt/services"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.LoadHTMLFiles()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.LoadHTMLGlob("views/*.html")


	router.GET("/", func(c *gin.Context) {
		go services.RunTimer()
		data, err := ioutil.ReadFile("status.json")

		if err != nil {
			log.Panic("failed reading data from file")
			return
		}
		var (
			info models.Info
		)

		err = json.Unmarshal(data, &info)

		if err != nil {
			log.Panic("failed to decode json")
			return
		}
		windIndicator := helpers.GenerateIndicator("wind", int(info.Status.Wind))
		waterIndicator := helpers.GenerateIndicator("water", int(info.Status.Water))
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"wind": info.Status.Wind,
			"water": info.Status.Water,
			"wind_cond": windIndicator,
			"water_cond": waterIndicator,
		})
	})
	router.Run(":4000")
}
