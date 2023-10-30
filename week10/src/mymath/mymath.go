package mymath

// func myAbs(number int) int // 외부 파일에 함수를 공개하기 위해선 함수의 첫글자를 대문자로 해준다.
func MyAbs(number int) int {
	if number < 0 {
		return number * -1
	} else {
		return number
	}

}

func MyPower(num1 int, num2 int) int {
	result := 1
	for i := 1; i <= num2; i++ {
		result = result * num1
	}
	return result
}
