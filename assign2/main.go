package main

import (
	"fmt"
	go_bmi "github.com/armstrongli/go-bmi"
	"strconv"
)

const noOfIndex int = 1

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	inputSet := [noOfIndex]map[string]string{}
	var totalFatRate float64 = 0.0

	fmt.Println("=========================================")
	fmt.Println("Welcome to Fat Rate Calculation Program")
	fmt.Println("=========================================")
	fmt.Println("We will calculate for ", noOfIndex, " persons in this program.")
	fmt.Println()
	//var sexFactor float64 = 0.0
	for i := 0; i < noOfIndex; i++ {

		fmt.Println("Person ", i, " :")
		fmt.Println("-----------------")
		name, weight, tall, age, sex := getInfo()
		//bmi := weight / (tall * tall)
		bmi, error := go_bmi.BMI(weight, tall)
		if error != nil {
			panic(error)
		}

		//var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexFactor)) / 100
		fatRate, badyFatCat, error := go_bmi.CalFatRate(bmi, age, sex)
		if error != nil {
			panic(error)
		}

		totalFatRate += fatRate

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
