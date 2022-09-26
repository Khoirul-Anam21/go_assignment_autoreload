package services

import (
	"fmt"
	"go_assignment_jwt/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func TestFileAfterUpdate(filepath string) {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Panic("failed reading data from file")
		return
	}
	var (
		model models.Info
	)

	err = json.Unmarshal(data, &model)

	if err != nil {
		log.Panic("failed to decode json")
		return
	}

	fmt.Println("Hasil baca file ", filepath, " -> ", model)

}