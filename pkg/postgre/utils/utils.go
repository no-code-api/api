package utils

import (
	"fmt"
)

func SqlString(value any) string {
	a := fmt.Sprintf("'%v'", value)
	return a
}
