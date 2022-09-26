package services

import (
	"encoding/json"
	"fmt"
	"go_assignment_jwt/helpers"
	"go_assignment_jwt/models"
	// "net/http"
	// "sync"

	// "io"
	"io/ioutil"
	"log"

	// "github.com/gin-gonic/gin"
)

func UpdateFileWithNewValue() {
	fmt.Println("membuka file...")
	data, err := ioutil.ReadFile("./status.json");

	if err != nil {
		log.Panic("failed reading data from file");
		return
	}
	var (
		info models.Info
		newWind int = helpers.GenerateRandomNum(1, 100)
		newWater int = helpers.GenerateRandomNum(1, 100);
	)
	

	err = json.Unmarshal(data, &info);

	if err != nil {
		log.Panic("failed to decode json");
		return
	}

	fmt.Println("Nilai air baru: ", newWater);
	fmt.Println("Nilai angin baru: ", newWind);

	info.Status.Water = int64(newWater)
	info.Status.Wind = int64(newWind)

	jsonString, err := json.Marshal(info);

	if err != nil {
		log.Panic("failed to encode json");
		return
	}

	err = ioutil.WriteFile("./status.json", jsonString, 0644)

	if err != nil {
		log.Panic("failed update json file");
		return
	}

	// c.HTML(http.StatusOK, "index.html", info.Status)
	
	fmt.Println(info);
}