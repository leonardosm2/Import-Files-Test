package main

import "fmt"

func divide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("erro")
		}
	}()

	result = a / b

	return result, err
}

func main() {
	for i := 10; i >= 0; i-- {
		r, err := divide(10, i)
		fmt.Printf("%d / %d = %d\n", 10, i, r)
		if err != nil {
			fmt.Printf("erro de cálculo: %v\n", err)
		}
	}
}
