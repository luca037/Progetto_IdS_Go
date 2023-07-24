package sources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	ApiKey string
}

const kMaxPages = 5  // Numero di pagine totali richieste alle api.
const kDimPage = 200 // Dimensione pagina richiesta alle api.

// Il numero di articli scaricati è kMaxPages * kDimPage.
func (guardian *Guardian) Download() []Article {
	// contenuti delle risposte alle chiamate api
	responsesBytes := make(chan []byte, kMaxPages)

	url := "https://content.guardianapis.com/search?show-fields=all&page-size=200&api-key="
	url += guardian.ApiKey
	go getResponses(responsesBytes, url, kMaxPages)

	// lista in cui salvo tutti i 1000 articoli delle risposte
	allArticles := make([]Article, kDimPage*kMaxPages)
	index := 0

	for content := range responsesBytes {
		// oggetto in cui salvo la risposta
		var response ResponseWrapper

		if err := json.Unmarshal(content, &response); err != nil {
			panic(err)
		}

		// salvo tutti gli articoli nella lista
		for _, article := range response.Response.Results {
			allArticles[index] = article
			index++
		}
	}

	return allArticles
}

// Permette di ottenere il contenuto delle chiamate alle api.
// ch è il canale in cui vengono inserire i contenuti delle risposte.
// url è l'indirizzo.
// nPages è il numero di pagine da scaricare.
func getResponses(ch chan<- []byte, url string, nPages int) {
	var senders sync.WaitGroup
	defer close(ch)

	for i := 1; i <= nPages; i++ {
		senders.Add(1)
		go func(n int) {
			url += "&page=" + fmt.Sprint(n)
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			var buffer bytes.Buffer
			io.Copy(&buffer, resp.Body)

			ch <- buffer.Bytes()
			senders.Done()
		}(i) // RICORDARSI DI MODIFICARE: per i test setta i=1
	}
	senders.Wait()
}

// Permette di evitare di creare l'oggetto fields presente nelle rieposte
// json del The Guardian.
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
