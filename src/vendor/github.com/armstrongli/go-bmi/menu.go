package gobmi

import "fmt"

func BodyFatCalculator() {
	defer fmt.Printf("欢迎下次使用BMI体脂计算器.\n\n")

	var (
		name, sex      string
		height, weight float64
		sex_code, age  int
	)

Menu:
	fmt.Println("**** BMI体脂计算器 ****")
	fmt.Scanf("姓名: %s\n", &name)
	fmt.Scanf("性别(男/女): %s\n", &sex)
	fmt.Scanf("身高(米): %.2f\n", &height)
	fmt.Scanf("体重(公斤): %.2f\n", &weight)
	fmt.Scanf("年龄(岁): %d\n", &age)

	sex_code = 1
	// TODO: 判断输入参数是否正确, 如果错误, 则goto Menu展示继续输入的界面
	switch {
	case height <= 0.4 || height >= 2.40:
		fmt.Println("输入的身高不合法")
		goto Menu
	case weight <= 7 || weight >= 120:
		fmt.Println("输入的体重不合法")
		goto Menu
	case age < 0 || age >= 120:
		fmt.Println("输入的年龄不合法")
		goto Menu
	case sex == "女":
		sex_code = 0
	}

	bmi := BMI(weight, height)
	fat := Fat(bmi, age, sex_code)
	Distribution(sex_code, age, fat)
}
