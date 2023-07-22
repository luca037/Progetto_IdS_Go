package patterns

import (
	"sort"
	"strings"
	"../sources"
)

type FrequencyPerArticleStrategy struct { }

func (strategy *FrequencyPerArticleStrategy) Execute(articles []sources.Article) []struct{Key string; Value int} {
    // mappa che funge da memoria
    memory := make(map[string]int)
    
    for _, article := range articles {
        fullText := append(strings.Split(article.Title, " "), strings.Split(article.Body, " ")...)

        // inserisico tutte le parole in un set per rimuovere i doppioni
        set := map[string]struct{}{}
        for _, word := range fullText {
            set[word] = struct{}{}
        }

        // effettuo il conteggio
        for word := range set {
            v := memory[word]
            memory[word] = v+1
        }
    }

    // slice in cui savo i risultati finali
    var sorted []struct{
        Key string
        Value int
    }

    // inserisico tutte le entry della mappa nello slice
    for k, v := range memory {
        sorted = append(sorted, struct {
            Key string
            Value int
        }{Key: k, Value: v})
    }

    // riordino in ordine crescente
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].Value > sorted[j].Value
    })

    return sorted
}
