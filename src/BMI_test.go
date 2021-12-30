package main

import (
	"testing"

	gobmi "github.com/armstrongli/go-bmi"
)

func TestBMI(t *testing.T) {
	height, weight := 1.78, 60.0
	bmi := gobmi.BMI(weight, height)
	t.Logf("bmi为: %.2f", bmi)
}

func TestFat(t *testing.T) {
	sex_code, age, height, weight := 1, 23, 1.78, 60.0
	bmi := gobmi.BMI(weight, height)
	fat := gobmi.Fat(bmi, age, sex_code)
	t.Logf("fat体脂率为: %.2f", fat)
}

func TestAll(t *testing.T) {
	sex_code, age, height, weight := 1, 23, 1.78, 60.0
	bmi := gobmi.BMI(weight, height)
	fat := gobmi.Fat(bmi, age, sex_code)
	gobmi.Distribution(sex_code, age, fat)
}
