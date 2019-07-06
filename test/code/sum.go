package code

import "fmt"

func Div(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("y can not zero")
	}
	return x / y, nil
}
