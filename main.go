package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/luca037/Progetto_Ids_Go/patterns"
	"github.com/luca037/Progetto_Ids_Go/sources"
)

func main() {
	// ### GESTIONE COMANDI UTENTE ###
	// opzioni disponibili
	dowload := flag.Bool("d", false, "Dowload articles")
	extract := flag.Bool("e", false, "Exctract terms")
	dowloadAndExtract := flag.Bool("de", false, "Dowload then extract terms")

	apiKey := flag.String("ak", "noApi", "Set API-KEY")

	help := flag.Bool("h", false, "Print this help message")

	flag.Parse()

	// comando help
	if *help {
		flag.PrintDefaults()
		return
	}

	// serializzatore/deserializzatore serve in entrambe le fasi
	var serializer sources.XmlSerializer = sources.XmlSerializer{
		DirectoryPath: "outputXml/",
	}

	// ### FASE DI DOWLOAD ###
	if *dowload || *dowloadAndExtract {
		if *apiKey == "noApi" {
			log.Println("ERROR - You need to set API-KEY")
			return
		}

		var factory patterns.SourceFactory
		var guardian sources.Source = factory.CreateSource(
			"Guardian",
			*apiKey,
		)
		var nytimes sources.Source = factory.CreateSource(
			"NYTimes",
			"./nytimes_articles_v2.csv",
		)

		log.Println("INFO - Starting download")
		allArticles := append(guardian.Download(), nytimes.Download()...)

		log.Println("INFO - Serializing articles")
		serializer.Serialize(allArticles)
	}

	// ### FASE DI ESTRAZIONE ###
	if *extract || *dowloadAndExtract {
		log.Println("INFO - Deserializing articles")
		deserializedArticles := serializer.Deserialize()

		var counter patterns.WordCounter = patterns.WordCounter{
			Strategy: &patterns.FrequencyPerArticleStrategy{},
		}

		log.Println("INFO - Extracting terms")
		results := counter.Count(deserializedArticles)

		// stampa primi 50 risultati
		outputFile, err := os.OpenFile("results.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		for i, entry := range results {
			if i == 50 {
				break
			}
			line := entry.Key + " " + fmt.Sprint(entry.Value) + "\n"
			outputFile.WriteString(line)
		}

		log.Println("INFO - You can find the results in", outputFile.Name())
	}
}
