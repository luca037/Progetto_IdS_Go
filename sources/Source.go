package sources

type Source interface {
	Download() []Article
}
