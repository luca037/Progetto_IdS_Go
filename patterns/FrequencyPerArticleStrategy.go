package patterns

import (
	"sort"
	"strings"
	"unicode"

	"github.com/luca037/Progetto_Ids_Go/sources"
)

type FrequencyPerArticleStrategy struct{}

func (strategy *FrequencyPerArticleStrategy) Execute(articles []sources.Article) []struct {
	Key   string
	Value int
} {
	// mappa che funge da memoria
	memory := make(map[string]int)

	for _, article := range articles {
		// rimovo punteggiatura dal titolo e corpo
		title := removePunctuationAndToLower(article.Title)
		body := removePunctuationAndToLower(article.Body)

		// unisco titolo e corpo in un'unico slice
		fullText := append(strings.Split(title, " "), strings.Split(body, " ")...)

		// inserisico tutte le parole in un set per rimuovere i doppioni
		set := map[string]struct{}{}
		for _, word := range fullText {
			set[word] = struct{}{}
		}

		// effettuo il conteggio
		for word := range set {
			if len(word) == 0 { // gestione carattere fantasma
				continue
			}
			v := memory[word]
			memory[word] = v + 1
		}
	}

	// slice in cui savo i risultati finali
	var sorted []struct {
		Key   string
		Value int
	}

	// inserisico tutte le entry della mappa nello slice
	for k, v := range memory {
		sorted = append(sorted, struct {
			Key   string
			Value int
		}{Key: k, Value: v})
	}

	// riordino in ordine crescente
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}

// Rimuove la punteggiatura da una stringa e la converte in minuscolo
func removePunctuationAndToLower(input string) string {
	var result strings.Builder

	for _, r := range input {
		if !unicode.IsPunct(r) {
			result.WriteRune(r)
		}
	}

	return strings.ToLower(result.String())
}
