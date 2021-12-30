package go_bmi

import "fmt"

func BMI(weightKG, heightM float64) (bmi float64, err error) {
	fmt.Println("This is replace module.")
	if weightKG <= 0 {
		err = fmt.Errorf("weight cannot be zero or negative")
		return
	}
	if heightM <= 0 {
		err = fmt.Errorf("height cannot be zero or negative")
		return
	}

	bmi = weightKG / (heightM * heightM)
	return
}
