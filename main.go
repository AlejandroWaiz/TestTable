package main

import "log"

func main() {

	var (
		FirstOperand  = 10
		SecondOperand = 0
	)

	result, err := Divide(FirstOperand, SecondOperand)

	if err != nil {
		log.Println(result, err)
	}

	log.Println(result, err)
}

type ErrDivideByZero struct{}

func (err *ErrDivideByZero) Error() string {
	return "Error!: Cant divide by zero, you want to destroy the world??"
}

func Divide(firstValue, secondValue int) (int, error) {

	if secondValue == 0 {
		return 0, &ErrDivideByZero{}
	}

	result := firstValue / secondValue

	return result, nil

}
