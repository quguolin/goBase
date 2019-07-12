package div

import "fmt"

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("分母不能为0")
	}
	return a / b, nil
}
