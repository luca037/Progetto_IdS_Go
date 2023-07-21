package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type ResponseWrapper struct {
	Response Response
}

type Response struct {
	Status   string
	PageSize int
	Results  []Article
}

type Guardian struct {
    OutputDir string 
    ApiKey  string
}

func (guardian *Guardian) Download() []Article {
    // creo la cartella in cui salvare i file (se non esiste)
    if err := os.MkdirAll(guardian.OutputDir, os.ModePerm); err != nil {
        panic(err)
    }

    // scarico le risposte json
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        // comando da eseguire
        cmd := join(
            "echo > ", guardian.OutputDir, "/", fmt.Sprint(i+1), ".json ",
            " && curl -o ", guardian.OutputDir, "/" ,fmt.Sprint(i+1), ".json ",
            "\"https://content.guardianapis.com/search?show-fields=all&page-size=200&", 
            "page=", fmt.Sprint(i+1),
            "&api-key=", guardian.ApiKey, "\"",
        )

        // eseguo il comando
        wg.Add(1)
        go executeBashCmd(&wg, cmd)
    }
    wg.Wait()

    // estraggo gli articoli dai file
    jsonFiles, err := os.ReadDir(guardian.OutputDir)
    if err != nil {
        panic(err)
    }

    // lista in cui salvo tutti i 1000 articoli delle risposte
    allArticles := make([]Article, 1000)
    index := 0

    for _, file := range jsonFiles {
        // risalgo al percorso assoluto file json e salvo il contenuto 
        filePath := join(guardian.OutputDir, "/", file.Name())
        content, _ := ioutil.ReadFile(filePath)

        // oggetto in cui salvo la risposta
        var response ResponseWrapper 

        err := json.Unmarshal(content, &response)
        if err != nil {
            panic(err)
        }

        // salvo tutti gli articoli nella risposta
        for _, article := range response.Response.Results {
            allArticles[index] = article
            index++
        }
    }

    return allArticles
}

// Permette di concatenare stringhe.
// strs è l'array di stringhe da concatenare.
func join(strs ...string) string {
    var builder strings.Builder
    for _, str := range strs {
        builder.WriteString(str)
    }
    return builder.String()
}

// Permette di eseguire un comando bash
func executeBashCmd(wg *sync.WaitGroup, cmd string) {
    execCmd := exec.Command("bash", "-c", cmd)
    if err := execCmd.Run(); err != nil {
        panic(err)
    }
    wg.Done()
}

func (a *Article) UnmarshalJSON(data []byte) error {
	type Alias Article // Crea un alias per evitare ricorsione infinita
	aux := &struct {
		Fields map[string]string `json:"fields"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Estrarre il valore "bodyText" dalla mappa "Fields"
	if bodyText, ok := aux.Fields["bodyText"]; ok {
		a.Body = bodyText
	}

	return nil
}
