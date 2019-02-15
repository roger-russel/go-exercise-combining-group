package main

import (
	"fmt"

	"github.com/roger-russel/go-exercise-combining-group/internal/pkg/csv"
	"github.com/roger-russel/go-exercise-combining-group/internal/pkg/pareto"
)

func main() {

	csvReader := csv.Reader("./assets/groups.csv")
	groups := pareto.Csv(csvReader)
	fmt.Printf("%v", groups)

}
