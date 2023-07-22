package sources

import (
	"encoding/xml"
	"fmt"
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
