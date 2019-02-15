package csv

import (
	"bufio"
	"encoding/csv"
	"os"
)

func Reader(fullFileName string) *csv.Reader {

	csvFile, _ := os.Open(fullFileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	return reader

}
