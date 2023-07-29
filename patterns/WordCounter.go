package patterns

import (
	"github.com/luca037/Progetto_Ids_Go/sources"
)

type WordCounter struct {
	Strategy WordCountStrategy
}

func (counter *WordCounter) Count(articles []sources.Article) []struct {
	Key   string
	Value int
} {
	return counter.Strategy.Execute(articles)
}
