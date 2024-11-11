package utils

import (
	"fmt"
)

func SqlString(value any) string {
	a := fmt.Sprintf("'%v'", value)
	fmt.Println("Value: ", value)
	fmt.Println("A: ", a)
	return a
}
