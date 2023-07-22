package sources

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type XmlSerializer struct {
	DirectoryPath string
}

func (serializer *XmlSerializer) Serialize(articles []Article) {
	// creo la cartella se non esiste
	os.MkdirAll(serializer.DirectoryPath, os.ModePerm)

	for i, article := range articles {
		filePath := serializer.DirectoryPath + fmt.Sprint(i+1) + ".xml"

		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		encoder := xml.NewEncoder(file)
		err = encoder.Encode(article)
		if err != nil {
			panic(err)
		}
	}
}

func (serializer *XmlSerializer) Deserialize() []Article {
	xmlFiles, err := ioutil.ReadDir(serializer.DirectoryPath)
	if err != nil {
		panic(err)
	}

	// slice in cui salvo tutti gli articoli deserializzati
	articles := make([]Article, len(xmlFiles))
	index := 0

	for _, fileInfo := range xmlFiles {
		filePath := serializer.DirectoryPath + fileInfo.Name()

		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		decoder := xml.NewDecoder(file)
		err = decoder.Decode(&articles[index])
		if err != nil {
			panic(err)
		}
		index++
	}

	return articles
}
