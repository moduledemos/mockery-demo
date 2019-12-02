package demo

type Story interface {
	ReadChapter(chapter int) ([]byte, error)
	HasChapter(chapter int) bool
}
