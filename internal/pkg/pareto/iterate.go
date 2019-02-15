package pareto

import "fmt"

//Par is a exploded possibilities from a group
type Par struct {
	left  string
	right string
	count int32
}

var Pares *[]Par

func Combine(line *[]string) {

	for index, element := range *line {

		if element == "" {
			continue
		}

		fmt.Println(index, element)
	}

}
