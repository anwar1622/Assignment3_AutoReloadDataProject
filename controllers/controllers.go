package controllers

import (
	"Auto_Reload/models"
	"Auto_Reload/services"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func AutoReload() {
	for {
		water := services.RandomNumberWater()
		wind := services.RandomNumberWind()

		num := models.StatusWaterWind{}
		num.Status.Water = water
		num.Status.Wind = wind

		jsonData, err := json.Marshal(num)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = ioutil.WriteFile("./data.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}

func ReloadWeb(write http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var status models.StatusWaterWind

	err = json.Unmarshal(jsonData, &status)
	if err != nil {
		log.Fatal(err.Error())
	}

	water := status.Status.Water
	wind := status.Status.Wind

	var (
		statusWater string
		statusWind  string
	)

	statusWater = services.Water(water)
	statusWind = services.Wind(wind)

	data := map[string]interface{}{
		"statusWater": statusWater,
		"statusWind":  statusWind,
		"water":       water,
		"wind":        wind,
	}
	fmt.Println("Water :", data["water"])
	fmt.Println("StatusWater :", data["statusWater"])
	fmt.Println("Wind :", data["wind"])
	fmt.Println("StatusWind :", data["statusWind"])

	template, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	template.Execute(write, data)
}
