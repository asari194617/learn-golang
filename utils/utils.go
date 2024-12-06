package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// 渡されたクエリパラメーターの値をもとに合計を計算
func Calculate(expression string) (int, error) {
	if !strings.Contains(expression, " ") {
		num, err := strconv.Atoi(expression)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", expression)
		}
		return num, nil
	}

	parts := strings.Split(expression, " ")
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid expression: %s", expression)
	}

	sum := 0
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", part)
		}
		sum += num
	}

	return sum, nil
}
