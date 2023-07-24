package sources

import (
	"encoding/csv"
	"io"
	"os"
	"sync"
)

type NYTimes struct {
	CsvInput string
}

const kArticlesPerFile = 1001 // Numero di articoli presenti nel file csv
const kCount = 1              // Quante volte estraggo gli articoli dal file csv (per test)

func (nytimes *NYTimes) Download() []Article {
	// canale in cui vengono inviati gli articoli estratti dal file csv
	articlesCh := make(chan Article, kArticlesPerFile*kCount)
	go extractArticles(articlesCh, nytimes.CsvInput)

	// lista in cui salvo tutti gli articoli presenti nel file
	articles := make([]Article, kArticlesPerFile*kCount) // il numero di articoli totali lo conosco già
	index := 0

	for article := range articlesCh {
		articles[index] = article
		index++
	}

	return articles
}

func extractArticles(articlesCh chan<- Article, filePath string) {
	var senders sync.WaitGroup

	for i := 0; i < kCount; i++ {
		senders.Add(1)

		go func() {
			// apro il file csv
			file, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			csvReader := csv.NewReader(file)

			// leggo tutte le righe del file csv
			for {
				rec, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					continue
				}

				// inserisco l'articolo nel canale
				articlesCh <- Article{Title: rec[2], Body: rec[3]}
			}
			senders.Done()
		}()
	}
	senders.Wait()
	close(articlesCh)
}

// Versione progetto originale del metodo dowload()
// func (nytimes *NYTimes) Download() []Article {
// 	file, err := os.Open(nytimes.CsvInput)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
//
// 	csvReader := csv.NewReader(file)
//
// 	// lista in cui salvo tutti gli articoli presenti nel file
// 	articles := make([]Article, 1001) // il numero di articoli totali lo conosco già
// 	index := 0
//
// 	// leggo tutte le righe del file csv
// 	for {
// 		rec, err := csvReader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			continue
// 		}
//
// 		articles[index] = Article{Title: rec[2], Body: rec[3]}
// 		index++
// 	}
//
// 	return articles
// }
