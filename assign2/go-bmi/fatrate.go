package go_bmi

import (
	"assignmnet/learngoassignment/assign2/util"
	"fmt"
)

var bodyFatMatrix = map[string]map[string][4]float64{
	util.Male: {
		util.AgeRange1: [4]float64{0.10, 0.16, 0.21, 0.26},
		util.AgeRange2: [4]float64{0.11, 0.17, 0.22, 0.27},
		util.AgeRange3: [4]float64{0.13, 0.19, 0.24, 0.29},
	},
	util.Female: {
		util.AgeRange1: [4]float64{0.20, 0.27, 0.34, 0.39},
		util.AgeRange2: [4]float64{0.21, 0.28, 0.35, 0.40},
		util.AgeRange3: [4]float64{0.22, 0.29, 0.36, 0.41},
	},
}

func CalFatRate(bmi float64, age int, sex string) (fatRate float64, badyFatCat string, err error) {
	if bmi <= 0 {
		err = fmt.Errorf("bmi must be greater than zero.")
		return
	}

	if age <= 0 || age > 150 {
		err = fmt.Errorf("Age must be greater than 0 and less than 150")
		return
	}

	if sex != util.Female && sex != util.Male {
		err = fmt.Errorf("Sex must m or f only.")
		return
	}

	var sexFactor float64
	if sex == util.Female {
		sexFactor = 0
	} else {
		sexFactor = 1
	}
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexFactor)) / 100
	badyFatCat = accessBodyFatStatus(sex, fatRate, age)
	return
}

func getBodyFatRangeCat(bodyFatMatrix [4]float64, fatrate float64) (cat string) {

	switch {
	case fatrate <= bodyFatMatrix[0]:
		return util.UnderFat
	case fatrate > bodyFatMatrix[0] && fatrate <= bodyFatMatrix[1]:
		return util.StandardMinus
	case fatrate > bodyFatMatrix[1] && fatrate <= bodyFatMatrix[2]:
		return util.StandardPlus
	case fatrate > bodyFatMatrix[2] && fatrate <= bodyFatMatrix[3]:
		return util.OverFat
	case fatrate > bodyFatMatrix[3]:
		return util.Obese
	}
	return

}

func getAgeRange(age int) (ageRange string) {
	switch {
	case age >= 18 && age <= 39:
		return util.AgeRange1
	case age >= 40 && age <= 59:
		return util.AgeRange2
	case age >= 60:
		return util.AgeRange3
	default:
		return util.AgeNotSupported
	}
}

func accessBodyFatStatus(sex string, fatrate float64, age int) (status string) {

	ageRange := getAgeRange(age)
	if ageRange == util.AgeNotSupported {
		return util.AgeNotSupported
	}

	return getBodyFatRangeCat(bodyFatMatrix[sex][ageRange], fatrate)

}
