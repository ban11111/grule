package grule

import (
	"fmt"
	"strings"
)

func GetTestRules(n int) (rawJsonRules string) {
	var list []string
	template := `{"name":"new%d", "value":%d, "comparator":"eq"}`
	for i := 0; i < n; i++ {
		fmt.Sprintf(template, i, i)
	}
	return "["+strings.Join(list, ",")+"]"
}