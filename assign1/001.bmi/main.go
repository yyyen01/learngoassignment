package main

import (
	"fmt"
	"strconv"
)

const noOfIndex int = 1
const female string = "f"
const male string = "m"
const underFat string = "Underfat"
const standardMinus string = "Standard Minus"
const standardPlus string = "Standard Plus"
const overFat string = "Over Fat"
const obese string = "Obese"
const ageRange1 string = "18-39"
const ageRange2 string = "40-59"
const ageRange3 string = "60+"
const ageNotSupported string = "Age not supported"

var bodyFatMatrix = map[string]map[string][4]float64{
	male: {
		ageRange1: [4]float64{0.10, 0.16, 0.21, 0.26},
		ageRange2: [4]float64{0.11, 0.17, 0.22, 0.27},
		ageRange3: [4]float64{0.13, 0.19, 0.24, 0.29},
	},
	female: {
		ageRange1: [4]float64{0.20, 0.27, 0.34, 0.39},
		ageRange2: [4]float64{0.21, 0.28, 0.35, 0.40},
		ageRange3: [4]float64{0.22, 0.29, 0.36, 0.41},
	},
}

func main() {
	inputSet := [noOfIndex]map[string]string{}
	var totalFatRate float64 = 0.0

	fmt.Println("=========================================")
	fmt.Println("Welcome to Fat Rate Calculation Program")
	fmt.Println("=========================================")
	fmt.Println("We will calculate for ", noOfIndex, " persons in this program.")
	fmt.Println()
	var sexFactor float64 = 0.0
	for i := 1; i <= noOfIndex; i++ {

		fmt.Println("Person ", i, " :")
		fmt.Println("-----------------")
		name, weight, tall, age, sex := getInfo()
		bmi := weight / (tall * tall)

		if sex == female {
			sexFactor = 0
		} else {
			sexFactor = 1
		}
		var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexFactor)) / 100
		totalFatRate += fatRate
		badyFatCat := accessBodyFatStatus(sex, fatRate, age)

		inputSet[i] = map[string]string{"name": name, "sex": sex, "age": strconv.Itoa(age), "weight": fmt.Sprintf("%.2f", weight), "height": fmt.Sprintf("%.2f", tall), "bmi": fmt.Sprintf("%.2f", bmi), "fatRate": fmt.Sprintf("%.2f", fatRate), "bodyFatCat": badyFatCat}
		fmt.Println("-----------------")
	}

	fmt.Println("=========================================")
	fmt.Println("Assessment Report")
	fmt.Println("=========================================")
	fmt.Println("Name Sex Age Weight Height BMI FatRate BodyFatCategory")
	fmt.Println("--------------------------------------------------------")
	for _, inputMap := range inputSet {
		fmt.Printf("%s %s %s %s %s %s %s %s\n", inputMap["name"], inputMap["sex"], inputMap["age"], inputMap["weight"], inputMap["height"], inputMap["bmi"], inputMap["fatRate"], inputMap["bodyFatCat"])
	}
	fmt.Println("=========================================")
	fmt.Printf("Average Fat Range for all people : %.2f", totalFatRate/float64(noOfIndex))

}

func getBodyFatRangeCat(bodyFatMatrix [4]float64, fatrate float64) (cat string) {

	switch {
	case fatrate <= bodyFatMatrix[0]:
		return underFat
	case fatrate > bodyFatMatrix[0] && fatrate <= bodyFatMatrix[1]:
		return standardMinus
	case fatrate > bodyFatMatrix[1] && fatrate <= bodyFatMatrix[2]:
		return standardPlus
	case fatrate > bodyFatMatrix[2] && fatrate <= bodyFatMatrix[3]:
		return overFat
	case fatrate > bodyFatMatrix[3]:
		return obese
	}
	return

}

func getAgeRange(age int) (ageRange string) {
	switch {
	case age >= 18 && age <= 39:
		return ageRange1
	case age >= 40 && age <= 59:
		return ageRange2
	case age >= 60:
		return ageRange3
	default:
		return ageNotSupported
	}
}

func accessBodyFatStatus(sex string, fatrate float64, age int) (status string) {

	ageRange := getAgeRange(age)
	if ageRange == ageNotSupported {
		return ageNotSupported
	}

	return getBodyFatRangeCat(bodyFatMatrix[sex][ageRange], fatrate)

}

func getInfo() (name string, weight float64, tall float64, age int, sex string) {

	fmt.Print("Name: ")
	fmt.Scanln(&name)

	var input string
	var err error

	for {
		fmt.Print("Age：")
		fmt.Scanln(&input)
		age, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid format. Please input a valid age number.")
		} else {
			break
		}
	}

	for {
		fmt.Print("Sex(m/f)：")
		fmt.Scanln(&input)
		if input != "m" && input != "f" {
			fmt.Println("Invalid format. Please input only m  or f.")
		} else {
			sex = input
			break
		}
	}

	for {
		fmt.Print("Weight (kg)：")
		fmt.Scanln(&input)
		weight, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid format. Please input a valid weight number(float only).")
		} else {
			break
		}
	}

	for {
		fmt.Print("Height (m)：")
		fmt.Scanln(&input)
		tall, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid format. Please input a valid height number(float only).")
		} else {
			break
		}
	}

	return
}
