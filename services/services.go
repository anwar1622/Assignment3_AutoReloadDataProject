package services

func Water(water int) string {
	var statusWater string
	if water <= 5 {
		statusWater = "Aman"
	} else if water >= 6 && water <=
		8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}
	return statusWater
}

func Wind(wind int) string {
	var statusWind string
	if wind <= 6 {
		statusWind = "Aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}
	return statusWind
}
