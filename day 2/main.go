package main

import "fmt"

func main() {
	var temp, finalTemp float32
	var temptype, finalType string
	temp, temptype = getTemp()
	finalTemp, finalType = convTemp(temp, temptype)
	if (finalType == "Invalid type!") {
		fmt.Println(finalType)
	} else {
		fmt.Println(fmt.Sprintf("The temperature is %f" + " degrees %v", finalTemp, finalType))
	}
}

func convTemp(temp float32, temptype string) (float32, string) {
	var finalTemp float32
	var finalType string
	if (temptype == "kelvin" || temptype == "Kelvin") {
		finalType = "Celsius."
		finalTemp = temp - 273.15
	} else if (temptype == "celsius" || temptype == "Celsius") {
		finalType = "Fahrenheit."
		finalTemp = (temp * 1.8) + 32
	} else if (temptype == "fahrenheit" || temptype == "Fahrenheit") {
		finalType = "Kelvin."
		finalTemp = (temp - 32) * 5/9 + 273.15
	} else {
		finalType = "Invalid type!"
	}
	return finalTemp, finalType
}

func getTemp() (float32, string) {
	var temp float32
	var tempType string
	fmt.Println("What is the temperature?")
	fmt.Scanln(&temp)
	fmt.Println("What type of reading is this?(kelvin, fahrenheit, celsius)")
	fmt.Scanln(&tempType)
	return temp, tempType
}