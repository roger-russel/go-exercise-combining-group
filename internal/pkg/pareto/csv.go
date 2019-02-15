package pareto

import (
	"encoding/csv"
	"io"
	"log"
)

func Csv(reader *csv.Reader) *[]Par {

	for {

		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		Combine(&line)

	}

	return Pares

}
