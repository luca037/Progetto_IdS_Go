package patterns

import (
	"../sources"
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
