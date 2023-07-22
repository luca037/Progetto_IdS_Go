package patterns

import(
    "../sources"
)

type WordCountStrategy interface {
    Execute([]sources.Article) []struct{Key string; Value int}
}
