package patterns

import (
	"github.com/luca037/Progetto_Ids_Go/sources"
)

type WordCountStrategy interface {
	Execute([]sources.Article) []struct {
		Key   string
		Value int
	}
}
