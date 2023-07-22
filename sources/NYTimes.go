package sources

import (
	"encoding/csv"
	"io"
	"os"
)

type NYTimes struct {
	CsvInput string
}

func (nytimes *NYTimes) Download() []Article {
	file, err := os.Open(nytimes.CsvInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	// lista in cui salvo tutti gli articoli presenti nel file
	articles := make([]Article, 1001) // il numero di articoli totali lo conosco già
	index := 0

	// leggo tutte le righe del file csv
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		articles[index] = Article{Title: rec[2], Body: rec[3]}
		index++
	}

	return articles
}
