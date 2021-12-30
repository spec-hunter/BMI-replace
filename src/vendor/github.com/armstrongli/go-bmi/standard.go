package gobmi

import "fmt"

func Distribution(sex_code int, age int, fat float64) {
	if sex_code == 1 {
		MaleBodyFat(age, fat)
	} else if sex_code == 0 {
		FemaleBodyFat(age, fat)
	}
}

func MaleBodyFat(age int, fat float64) {
	switch {
	case (age >= 18 && age <= 39):
		switch {
		case (fat > 0 && fat <= 10):
			fmt.Printf("偏瘦.\n")
		case (fat > 10 && fat <= 16):
			fmt.Printf("标准.\n")
		case (fat > 16 && fat <= 21):
			fmt.Printf("偏重.\n")
		case (fat > 21 && fat <= 26):
			fmt.Printf("肥胖.\n")
		case (fat > 26 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 40 && age <= 59):
		switch {
		case (fat > 0 && fat <= 11):
			fmt.Printf("偏瘦.\n")
		case (fat > 11 && fat <= 17):
			fmt.Printf("标准.\n")
		case (fat > 17 && fat <= 22):
			fmt.Printf("偏重.\n")
		case (fat > 22 && fat <= 27):
			fmt.Printf("肥胖.\n")
		case (fat > 27 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 60):
		switch {
		case (fat > 0 && fat <= 13):
			fmt.Printf("偏瘦.\n")
		case (fat > 13 && fat <= 19):
			fmt.Printf("标准.\n")
		case (fat > 19 && fat <= 24):
			fmt.Printf("偏重.\n")
		case (fat > 24 && fat <= 29):
			fmt.Printf("肥胖.\n")
		case (fat > 29 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	}
}

func FemaleBodyFat(age int, fat float64) {
	switch {
	case (age >= 18 && age <= 39):
		switch {
		case (fat > 0 && fat <= 20):
			fmt.Printf("偏瘦.\n")
		case (fat > 20 && fat <= 27):
			fmt.Printf("标准.\n")
		case (fat > 27 && fat <= 34):
			fmt.Printf("偏重.\n")
		case (fat > 34 && fat <= 39):
			fmt.Printf("肥胖.\n")
		case (fat > 39 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 40 && age <= 59):
		switch {
		case (fat > 0 && fat <= 21):
			fmt.Printf("偏瘦.\n")
		case (fat > 21 && fat <= 28):
			fmt.Printf("标准.\n")
		case (fat > 28 && fat <= 35):
			fmt.Printf("偏重.\n")
		case (fat > 35 && fat <= 40):
			fmt.Printf("肥胖.\n")
		case (fat > 40 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 60):
		switch {
		case (fat > 0 && fat <= 22):
			fmt.Printf("偏瘦.\n")
		case (fat > 22 && fat <= 29):
			fmt.Printf("标准.\n")
		case (fat > 29 && fat <= 36):
			fmt.Printf("偏重.\n")
		case (fat > 36 && fat <= 41):
			fmt.Printf("肥胖.\n")
		case (fat > 41 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	}
}
