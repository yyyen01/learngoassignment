package go_bmi

import (
	"assignmnet/learngoassignment/assign2/util"
	"fmt"
	"testing"
)

func TestCalFatRate1(t *testing.T) {
	_, _, err := CalFatRate(-1, 1, "0")
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestCalFatRate2(t *testing.T) {
	_, _, err := CalFatRate(1.1, -1, "0")
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestCalFatRate3(t *testing.T) {
	_, _, err := CalFatRate(1.1, 156, "0")
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestCalFatRate4(t *testing.T) {
	_, _, err := CalFatRate(1.1, 50, "0")
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestCalFatRate5(t *testing.T) {
	_, _, err := CalFatRate(1.1, 50, "0")
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestCalFatRate6(t *testing.T) {
	expFatRate := "0.22"
	expCat := util.StandardPlus
	bmi, _ := BMI(56, 1.8)
	fatrate, suggestion, err := CalFatRate(bmi, 76, "m")

	if err != nil {
		t.Errorf("should not be error, but err is returned.")
	}
	fatrate_str := fmt.Sprintf("%.2f", fatrate)
	if fatrate_str != expFatRate {
		t.Fatalf("fatrate should be %s, but return %s", expFatRate, fatrate_str)
	}
	if suggestion != expCat {
		t.Fatalf("Suggestion should be %s, but return %s", expCat, suggestion)
	}
}
